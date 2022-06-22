package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

var PORT int = 3000

var PAGE_URL string = "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html"

type Track struct {
	Description string            `json:"description"`
	CreatedAt   string            `json:"createdAt"`
	Formatted   map[string]string `json:"formatted"`
}

type TrackHistory struct {
	ReceivedBy string   `json:"receivedBy"`
	Histories  []*Track `json:"histories"`
}

type TrackingResponse struct {
	Status map[string]string `json:"status"`
	Data   *TrackHistory     `json:"data"`
}

var data TrackingResponse

func main() {
	scrapeDataTo(&data)

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/tracking", trackingHandler).Methods("GET")

	log.Println("Web application listening on port", PORT)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", PORT), router),
	)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Tracking Detail Scraper by Muhammad Abdis")
}

func trackingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

func scrapeDataTo(data *TrackingResponse) {
	res, err := http.Get(PAGE_URL)

	if err != nil {
		log.Println("Cannot reach requested URL")
	}

	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	receivedBy := doc.Find(".tracking:last-child tbody tr:last-child td:nth-child(2)").Text()
	re := regexp.MustCompile(`\[(.*?)\s+\|`)
	receivedBy = re.FindStringSubmatch(receivedBy)[1]

	histories := make([]*Track, 0)
	doc.Find(".tracking:last-child tbody tr").Each(func(i int, s *goquery.Selection) {
		createdAt, _ := time.Parse("02-01-2006 15:04 MST", s.Find("td:first-child").Text()+" WIB")

		monthReplacer := strings.NewReplacer(
			"January", "Januari",
			"February", "Februari",
			"March", "Maret",
			"May", "Mei",
			"April", "April",
			"June", "Juni",
			"July", "Juli",
			"August", "Agustus",
			"September", "September",
			"October", "Oktober",
			"November", "November",
			"December", "Desember",
		)

		histories = append(histories, &Track{
			Description: s.Find("td:nth-child(2)").Text(),
			CreatedAt:   createdAt.Format(time.RFC3339),
			Formatted: map[string]string{
				"createdAt": monthReplacer.Replace(
					createdAt.Format("02 January 2006, 15:04 MST"),
				),
			},
		})
	})

	data.Status = map[string]string{
		"code":    "060101",
		"message": "Delivery tracking detail fetched successfully",
	}
	data.Data = &TrackHistory{
		ReceivedBy: receivedBy,
		Histories:  histories,
	}
}
