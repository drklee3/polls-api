package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/drklee3/polls-api/api/app/model"
	"github.com/jinzhu/gorm"
)

// GetAllPolls gets the list of all polls
func GetAllPolls(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// array of pointers so the loop modifies instead of clones
	polls := []*model.Poll{}
	db.Find(&polls)

	for _, poll := range polls {
		// unmarshall each poll content
		if err := poll.UnmarshalContent(); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	respondJSON(w, http.StatusOK, polls)
}

// CreatePoll creates a new poll
func CreatePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	poll := model.Poll{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&poll); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// initialize poll and check err
	if err := poll.Initialize(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	for {
		poll.SetUUID()

		if !hasUUID(db, &poll) {
			break
		}
	}

	// marshal content / serialize back to json
	if err := poll.MarshalContent(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Save(&poll).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := poll.UnmarshalContent(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, poll)
}

// GetPoll gets a single poll
func GetPoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// get the poll item
	poll, err := getPoll(db, w, r, false)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	if err := poll.UnmarshalContent(); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, poll)
}

// VotePoll creates a poll submission
func VotePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	poll, err := getPoll(db, w, r, true)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	// unmarshal content / json to struct
	if err := poll.UnmarshalContent(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	submission := model.Submission{
		CreatedAt: time.Now(),
		IP:        strings.Split(r.RemoteAddr, ":")[0],
		PollID:    poll.ID,
	}

	// check if restricted to single vote
	if poll.Content.Options.Restrictions == "single" {
		// if submission exists, return with error
		if hasSubmission(db, &submission, poll) {
			respondError(w, http.StatusBadRequest, "user already voted")
			return
		}
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&submission); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// add submission
	if err := poll.AddSubmission(&submission); err != nil {
		// check for errors when adding submission
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// marshal content / serialize back to json
	if err := poll.MarshalContent(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// save submission
	if err := db.Save(&submission).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// save poll data
	if err := db.Save(&poll).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, poll)
}

// UpdatePoll updates poll options
func UpdatePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	poll, err := getPoll(db, w, r, false)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	var updatedPoll model.Poll
	// decode body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedPoll); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// update poll with new data without modifying counts
	poll.Update(&updatedPoll)

	// save new poll
	if err := db.Save(&poll).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, poll)
}

// DeletePoll deletes a single poll
func DeletePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	poll, err := getPoll(db, w, r, false)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	if err := db.Delete(&poll).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, nil)
}

// ArchivePoll disables further submissions for a single poll
func ArchivePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	poll, err := getPoll(db, w, r, false)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	poll.Archive()
	if err := db.Save(&poll).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, poll)
}

// RestorePoll re-enables submissions for a single poll
func RestorePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	poll, err := getPoll(db, w, r, false)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	poll.Restore()
	if err := db.Save(&poll).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, poll)
}
