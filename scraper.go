package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeTracking() {
	PAGE_URL := "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html"

	res, err := http.Get(PAGE_URL)

	if err != nil {
		log.Println("Cannot reach requested URL")
	}

	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	doc.Find(".tracking:last-child tbody tr").Each(func(i int, s *goquery.Selection) {
		createdAt, _ := time.Parse("02-01-2006 15:04 MST", s.Find("td:first-child").Text()+" WIB")

		tracking.Create(Track{
			Description: s.Find("td:nth-child(2)").Text(),
			CreatedAt:   createdAt,
		})
	})
}
