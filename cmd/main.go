package main

import (
	"fmt"
	"go-url-shortener/database"
	"go-url-shortener/usecase/shortenurl"
	"go-url-shortener/web"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("failed to load env file")
	}

	// connect to SQLite
	sqliteDB, err := database.ConnectSQLite()
	if err != nil {
		panic("failed to connect to sqliteDB")
	}

	database.RunMigrations(sqliteDB)

	// create the services with the db instance
	urlService := shortenurl.NewService(sqliteDB)

	// start the API Layer with the above services
	urlAPI := web.NewShortenURLAPI(urlService)

	// create the mux router
	router := mux.NewRouter()
	router.Handle("/url", http.HandlerFunc(urlAPI.CreateShortenURLHander)).Methods(http.MethodPost)
	router.Handle("/url", http.HandlerFunc(urlAPI.FetchOriginalURLHandler)).Methods(http.MethodGet)

	fmt.Println("Starting server")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), router); err != nil {
		log.Fatal("error starting the server")
	}
}
