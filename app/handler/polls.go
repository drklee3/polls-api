package handler

import (
	// "encoding/json"
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

func CreatePoll(w http.ResponseWriter, r *http.Request) {

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
