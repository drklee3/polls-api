package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drklee3/polls-api/app/model"
	"github.com/drklee3/polls-api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Dbname)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	// test connection
	_, err = db.Raw("SELECT 1 + 1 AS result").Rows()
	if err != nil {
		log.Fatalf("Failed to ping DB: %s", err)
	}

	// run migrations
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the polls
	a.Get("/polls", a.GetAllPolls)
	a.Post("/polls", a.CreatePoll)
	a.Get("/polls/{id:[0-9]+}", a.GetPoll)
	a.Put("/polls/{id:[0-9]+}", a.UpdatePoll)
	a.Delete("/polls/{id:[0-9]+}", a.DeletePoll)
	a.Put("/polls/{id:[0-9]+}/archive", a.ArchivePoll)
	a.Delete("/polls/{id:[0-9]+}/archive", a.RestorePoll)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/*
** Polls Handlers
 */
func (a *App) GetAllPolls(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllPolls")
	// handler.GetAllPolls(a.DB, w, r)
}

func (a *App) CreatePoll(w http.ResponseWriter, r *http.Request) {
	log.Println("CreatePoll")
	// handler.CreatePoll(a.DB, w, r)
}

func (a *App) GetPoll(w http.ResponseWriter, r *http.Request) {
	log.Println("GetPoll")
	// handler.GetPoll(a.DB, w, r)
}

func (a *App) UpdatePoll(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdatePoll")
	// handler.UpdatePoll(a.DB, w, r)
}

func (a *App) DeletePoll(w http.ResponseWriter, r *http.Request) {
	log.Println("DeletePoll")
	// handler.DeletePoll(a.DB, w, r)
}

func (a *App) ArchivePoll(w http.ResponseWriter, r *http.Request) {
	log.Println("ArchivePoll")
	// handler.ArchivePoll(a.DB, w, r)
}

func (a *App) RestorePoll(w http.ResponseWriter, r *http.Request) {
	log.Println("RestorePoll")
	// handler.RestorePoll(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Printf("Listening on %s", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
