package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"reverseProxy/internal/stats"
)

func main() {
	var port = flag.Int("p", 8081, "port")
	api := stats.NewStatsApi()
	log.Println("Started stats logger")
	http.HandleFunc("/stats", api.GetStats)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	fmt.Print(err)
}
