package main

import (
	"log"
	"net/http"

	"reverseProxy/internal/stats"
)

func main() {
	api := stats.NewStatsApi()
	log.Println("Started stats logger")
	http.HandleFunc("/stats", api.GetStats)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
