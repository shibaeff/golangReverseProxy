package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"reverseProxy/internal/stats"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	api := stats.NewStatsApi()
	log.Println("Started stats logger")
	http.HandleFunc("/stats", api.GetStats)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	fmt.Print(err)
}
