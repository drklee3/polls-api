package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
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
	gorm.Model
	Title    string `gorm:"not null"`
	Archived bool
	Content  postgres.Jsonb
}

// Archive archives a poll and disables submissions
func (p *Poll) Archive() {
	p.Archived = true
}

// Restore restores a poll and re-enables submissions
func (p *Poll) Restore() {
	p.Archived = false
}

// Submissions contain a single submission of
// a poll to keep track of duplicates
type Submissions struct {
	gorm.Model
	IP     string `gorm:"not null"`
	PollID uint
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Poll{}, &Submissions{})
	db.Model(&Submissions{}).AddForeignKey("poll_id", "polls(id)", "CASCADE", "CASCADE")
	return db
}
