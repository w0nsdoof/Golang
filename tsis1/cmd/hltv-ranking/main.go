package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routes")

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/players", Players).Methods("GET")
	router.HandleFunc("/players/{player}", Player).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
