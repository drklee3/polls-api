package handler

import (
	"encoding/json"
	"net/http"

	"github.com/drklee3/polls-api/app/model"
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

func GetPoll(w http.ResponseWriter, r *http.Request) {

}

func UpdatePoll(w http.ResponseWriter, r *http.Request) {

}

func DeletePoll(w http.ResponseWriter, r *http.Request) {

}

func ArchivePoll(w http.ResponseWriter, r *http.Request) {

}

func RestorePoll(w http.ResponseWriter, r *http.Request) {

}
