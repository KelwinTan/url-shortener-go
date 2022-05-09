package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/KelwinTan/url-shortener-go/config"
	"github.com/KelwinTan/url-shortener-go/db"
	"github.com/KelwinTan/url-shortener-go/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	var (
		err error
	)

	config := config.GetConfig()
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DBUsername, config.DBPassword, config.DBName)

	a.DB, err = sql.Open(config.DBDriver, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	DB := db.Init()
	dbHandler := handlers.New(DB)

	a.Router = mux.NewRouter()
	a.initializeRoutes(dbHandler)
}

func (a *App) initializeRoutes(dbHandler handlers.DBHandler) {
	//for testing ping pong
	a.Router.HandleFunc("/test", dbHandler.Default).Methods(http.MethodGet)

	a.Router.HandleFunc("/", dbHandler.GetUrls).Methods(http.MethodGet)
	a.Router.HandleFunc("/{short_url}", dbHandler.RedirectUrl).Methods(http.MethodGet)
	a.Router.HandleFunc("/{short_url}", dbHandler.DeleteUrl).Methods(http.MethodDelete)
	a.Router.HandleFunc("/{short_url}", dbHandler.UpdateUrl).Methods(http.MethodPut)
	a.Router.HandleFunc("/", dbHandler.ShortenURL).Methods(http.MethodPost)
}

func (a *App) Run() {
	log.Print("app running")
	config := config.GetConfig()
	log.Fatal(http.ListenAndServe(config.AppPort, a.Router))
}
