package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/drklee3/polls-api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllPolls gets the list of all polls
func GetAllPolls(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	polls := []model.Poll{}
	db.Find(&polls)
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

	if err := db.Save(&poll).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, poll)
}

// GetPoll gets a single poll
func GetPoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// get the poll item
	poll, err := getPoll(db, id, w, r)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}

	respondJSON(w, http.StatusOK, poll)
}

func UpdatePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func DeletePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func ArchivePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func RestorePoll(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func getPoll(db *gorm.DB, idStr string, w http.ResponseWriter, r *http.Request) (*model.Poll, error) {
	// parse id from string
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil, err
	}

	var poll model.Poll
	if err := db.First(&poll, model.Poll{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil, err
	}

	return &poll, nil
}
