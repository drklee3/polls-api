package handler

import (
	"database/sql"
	// "encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	// "github.com/gorilla/mux"
)

func GetAllPolls(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	/*
	rows, err := db.Query("SELECT 1 + 1 AS result")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(rows))
	*/
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
