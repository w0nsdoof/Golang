package main

import (
	"log"
	"net/http"

	"awesomeProject/api"
	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")

	api.SetupRoutes(router)

	http.Handle("/", router)

	log.Println("Server started on port", port)
	//start and listen to requests
	http.ListenAndServe(port, router)

}
