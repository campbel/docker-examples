package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health check called")
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":80", nil)
}
