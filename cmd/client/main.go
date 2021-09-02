package main

import (
	"log"
	"net/http"

	"reverseProxy/internal/stats"
)

func main() {
	api := stats.NewStatsApi()
	// port := 8081
	log.Println("Started stats logger")
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(fmt.Sprintf("hi on port %d", port)))
	//})
	http.HandleFunc("/stats", api.GetStats)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
