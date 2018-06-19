package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// PollChoice contains options for a single poll choice
type PollChoice struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Count uint   `json:"count"`
}

// PollOptions contains options for a poll
type PollOptions struct {
	Restrictions     string     `json:"restrictions"`
	PollType         string     `json:"poll_type"`
	RandomizeChoices bool       `json:"randomize_choices"`
	EndTime          *time.Time `json:"endtime"`
}

// PollContent contains questions / options for a poll
type PollContent struct {
	Choices []PollChoice `json:"choices"`
	Options PollOptions  `json:"opions"`
}

// Poll contains a single poll data
type Poll struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp"`
	UpdatedAt *time.Time
	Title     string      `gorm:"not null"`
	Archived  bool        `gorm:"not null"`
	Content   PollContent `gorm:"type:jsonb not null default '{}'::jsonb"`
}

// Archive archives a poll and disables submissions
func (p *Poll) Archive() {
	p.Archived = true
}

// Restore restores a poll and re-enables submissions
func (p *Poll) Restore() {
	p.Archived = false
}

// Value marshals data for jsonb column
func (p *PollContent) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Scan unmarshals data for jsonb column
func (p *PollContent) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i PollContent
	if err := json.Unmarshal(source, &i); err != nil {
		return err
	}

	return nil
}

// Submissions contain a single submission of
// a poll to keep track of duplicates
type Submissions struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	IP        string    `gorm:"not null"`
	PollID    uint64    `gorm:"not null"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Poll{}, &Submissions{})
	db.Model(&Submissions{}).AddForeignKey("poll_id", "polls(id)", "CASCADE", "CASCADE")
	return db
}
