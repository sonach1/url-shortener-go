package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"urlshortener/config"
	"urlshortener/handlers"
)

func main() {
	config.Init()
	r := mux.NewRouter()

	r.HandleFunc("/api/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortCode}", handlers.RedirectURL).Methods("GET")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
