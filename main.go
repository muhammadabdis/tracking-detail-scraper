package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var PORT int = 3000

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Tracking Detail Scraper by Muhammad Abdis")
	}).Methods("GET")

	log.Println("Web application listening on port", PORT)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", PORT), router),
	)
}
