package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("starting up")

	// create a /health endpoint, by default this will return status 200 with an empty body
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/health called")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
