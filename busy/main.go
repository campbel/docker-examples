package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting up...")

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				i := 0
				for {
					i++
				}
			}()
		}
	}()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Healthy")
	})
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}
}
