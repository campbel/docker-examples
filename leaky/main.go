package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("starting up...")

	go func() {
		var bigOleArray [][]string
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			bigOleArray = append(bigOleArray, make([]string, 65536))
		}
	}()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Healthy")
	})
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}
}
