package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var PORT int = 3000
var tracking Tracking

func main() {
	ScrapeTracking()

	router := mux.NewRouter()

	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/tracking", GetTracking).Methods("GET")

	log.Println("Web application listening on port", PORT)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", PORT), router),
	)
}
