package model

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
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
	Options PollOptions            `json:"options"`
}

// Poll contains a single poll data
type Poll struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp"`
	UpdatedAt *time.Time
	Title     string         `gorm:"not null"`
	Archived  bool           `gorm:"not null"`
	ContentB  postgres.Jsonb `sql:"type:jsonb" json:"-" gorm:"column:content; type:jsonb; not null; default '{}'::jsonb"`
	Content   PollContent    `sql:"-" gorm:"-" json:"content"`
}

// MarshalContent converts struct content to jsonb (serializes)
func (p *Poll) MarshalContent() error {
	var err error
	p.ContentB.RawMessage, err = json.Marshal(p.Content)
	return err
}

// UnmarshalContent converts json to a PollContent struct (deserializes)
func (p *Poll) UnmarshalContent() error {
	return json.Unmarshal(p.ContentB.RawMessage, &p.Content)
}

// Initialize sets initial values
func (p *Poll) Initialize() error {
	for key, val := range p.Content.Choices {
		val.Count = 0

		// parse key, set as id
		id, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			// invalid key
			return errors.New("invalid poll ID")
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

	// update updated time
	now := time.Now()
	p.UpdatedAt = &now
}

// AddSubmission adds a single submission to a poll
func (p *Poll) AddSubmission(s *Submission) error {
	// check for empty submission
	if len(s.ChoiceIDs) == 0 {
		return errors.New("submission cannot be empty")
	}

	// check if archived
	if p.Archived {
		return errors.New("poll is archived")
	}

	// check if radio type poll, remove all but first in submissions
	if p.Content.Options.PollType == "radio" {
		s.ChoiceIDs = s.ChoiceIDs[:1]
	}

	modified := 0

	for _, id := range s.ChoiceIDs {
		strID := strconv.FormatUint(uint64(id), 10)
		if val, ok := p.Content.Choices[strID]; ok {
			val.Count++
			modified++ // update modified choices
		}
	}

	// check if no choices were modified
	if modified == 0 {
		return errors.New("invalid submission choice")
	}

	return nil
}

// Submission contain a single submission of
// a poll to keep track of duplicates
type Submission struct {
	ID        uint64    `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	IP        string    `gorm:"not null"`
	PollID    uint64    `gorm:"not null"`
	ChoiceIDs []uint    `sql:"-" gorm:"-" json:"choice_ids"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Poll{}, &Submission{})
	db.Model(&Submission{}).AddForeignKey("poll_id", "polls(id)", "CASCADE", "CASCADE")
	return db
}
