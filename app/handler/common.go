package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/drklee3/polls-api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func parsePollID(r *http.Request) (uint64, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// parse id from string
	return strconv.ParseUint(idStr, 10, 64)
}

// getPoll gets a single poll by id and responds with 404 if not found
func getPoll(db *gorm.DB, w http.ResponseWriter, r *http.Request, shouldLock bool) (*model.Poll, error) {
	id, err := parsePollID(r)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil, err
	}

	// get poll from db
	var poll model.Poll

	// set for update row lock
	// this prevents data races in the db if a vote was made in the middle of another one
	// locks a single row for **updates**, allows reads still
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

// hasSubmission checks if a poll has a submission by a user (determined by ip)
func hasSubmission(db *gorm.DB, s *model.Submission, p *model.Poll) bool {
	return db.First(&s, model.Submission{IP: s.IP, PollID: p.ID}).Error == nil
}

// hasUUID checks if a poll has a given id
func hasUUID(db *gorm.DB, p *model.Poll) bool {
	return db.First(&p, model.Poll{UUID: p.UUID}).Error == nil
}
