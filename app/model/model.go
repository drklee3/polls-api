package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// PollChoice contains options for a single poll choice
type PollChoice struct {
	ID    uint64 `json:"id"`
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
	Choices map[string]*PollChoice `json:"choices"`
	Options PollOptions            `json:"opions"`
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

// Poll contains a single poll data
type Poll struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp"`
	UpdatedAt *time.Time
	Title     string      `gorm:"not null"`
	Archived  bool        `gorm:"not null"`
	Content   PollContent `gorm:"type:jsonb not null default '{}'::jsonb"`
}

// Initialize sets initial values
func (p *Poll) Initialize() error {
	for key, val := range p.Content.Choices {
		val.Count = 0

		// parse key, set as id
		id, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			// invalid key
			return err
		}
		val.ID = id
	}

	p.CreatedAt = time.Now()

	return nil
}

// Archive archives a poll and disables submissions
func (p *Poll) Archive() {
	p.Archived = true
}

// Restore restores a poll and re-enables submissions
func (p *Poll) Restore() {
	p.Archived = false
}

// Update updates a poll's options without modifying choice counts
func (p *Poll) Update(u *Poll) {
	for keyUpdated, valUpdated := range u.Content.Choices {
		// search previous choices for match
		previous, found := p.Content.Choices[keyUpdated]

		if found {
			// update existing choice to previous count
			u.Content.Choices[keyUpdated].Count = previous.Count
		} else {
			// (new choice) not found in previous poll
			valUpdated.Count = 0
			u.Content.Choices[keyUpdated] = valUpdated
		}
	}

	// set current poll data to new updated data
	// doesn't seem to work if using `p = u`? have to copy data over
	*p = *u
}

// AddSubmission adds a single submission to a poll
func (p *Poll) AddSubmission(s *SubmissionOptions) {
	for _, id := range s.ChoiceIDs {
		if val, ok := p.Content.Choices[string(id)]; ok {
			val.Count++
		}
	}
}

// Submission contain a single submission of
// a poll to keep track of duplicates
type Submission struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	IP        string    `gorm:"not null"`
	PollID    uint64    `gorm:"not null"`
}

// SubmissionOptions contains selected choice data from a submission
type SubmissionOptions struct {
	ChoiceIDs []uint `json:"choice_ids"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Poll{}, &Submission{})
	db.Model(&Submission{}).AddForeignKey("poll_id", "polls(id)", "CASCADE", "CASCADE")
	return db
}
