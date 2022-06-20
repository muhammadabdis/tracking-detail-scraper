package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var PORT int = 3000

type Track struct {
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	Formatted   struct {
		CreatedAt string `json:"createdAt"`
	} `json:"formatted"`
}

type TrackingResponse struct {
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data struct {
		ReceivedBy string   `json:"receivedBy"`
		Histories  []*Track `json:"histories"`
	} `json:"data"`
}

var baseData TrackingResponse

func main() {
	initTestData()

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/tracking", trackingHandler).Methods("GET")

	log.Println("Web application listening on port", PORT)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", PORT), router),
	)
}

func initTestData() {
	file, _ := os.Open("example.json")

	log.Println("Successfully read json data")

	defer file.Close()

	b, _ := io.ReadAll(file)
	json.Unmarshal(b, &baseData)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Tracking Detail Scraper by Muhammad Abdis")
}

func trackingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(baseData)
}
