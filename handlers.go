package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Tracking Detail Scraper by Muhammad Abdis")
}

func GetTracking(w http.ResponseWriter, r *http.Request) {
	tracks := tracking.Find()

	type History struct {
		Description string            `json:"description"`
		CreatedAt   string            `json:"createdAt"`
		Formatted   map[string]string `json:"formatted"`
	}

	type Data struct {
		ReceivedBy string     `json:"receivedBy"`
		Histories  []*History `json:"histories"`
	}

	type Response struct {
		Status map[string]string `json:"status"`
		Data   *Data             `json:"data"`
	}

	// Format for ReceivedBy
	receivedBy := tracks[len(tracks)-1].Description
	re := regexp.MustCompile(`\[(.*?)\s+\|`)
	receivedBy = re.FindStringSubmatch(receivedBy)[1]

	// Format for Histories
	var histories []*History

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

	for _, track := range tracks {
		histories = append(histories, &History{
			Description: track.Description,
			CreatedAt:   track.CreatedAt.Format(time.RFC3339),
			Formatted: map[string]string{
				"created_at": monthReplacer.Replace(
					track.CreatedAt.Format("02 January 2006, 15:04 MST"),
				),
			},
		})
	}

	// Format for Response
	response := Response{
		Status: map[string]string{
			"code":    "060101",
			"message": "Delivery tracking detail fetched successfully",
		},
		Data: &Data{
			ReceivedBy: receivedBy,
			Histories:  histories,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
