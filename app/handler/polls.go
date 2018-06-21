package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/drklee3/polls-api/app/model"
	"github.com/gorilla/mux"
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

	// marshal content / serialize back to json
	if err := poll.MarshalContent(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// initialize poll and check err
	if err := poll.Initialize(); err != nil {
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

	respondJSON(w, http.StatusOK, poll)
}

// VotePoll creates a poll submission
func VotePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	poll, err := getPoll(db, w, r, false)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	submission := model.Submission{
		CreatedAt: time.Now(),
		IP:        strings.Split(r.RemoteAddr, ":")[0],
		PollID:    poll.ID,
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&submission); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// unmarshal content / json to struct
	if err := poll.UnmarshalContent(); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// add submission
	if err := poll.AddSubmission(&submission); err != nil {
		// check for errors when adding submission
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// potential for data race here?
	// in the case another submission is made at this point
	// the new submission would have old value

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
	poll, err := getPoll(db, w, r, true)
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

// getPoll gets a single poll by id and responds with 404 if not found
func getPoll(db *gorm.DB, w http.ResponseWriter, r *http.Request, shouldLock bool) (*model.Poll, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// parse id from string
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil, err
	}

	// get poll from db
	var poll model.Poll

	// set for update row lock
	// https://www.postgresql.org/docs/current/static/explicit-locking.html#LOCKING-ROWS
	if shouldLock {
		db = db.Set("gorm:query_option", "FOR UPDATE")
	}

	if err := db.First(&poll, model.Poll{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil, err
	}

	return &poll, nil
}
