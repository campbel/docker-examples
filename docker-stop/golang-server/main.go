package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Waiting for docker stop signal")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world")
	})

	server := http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		<-sigs
		server.Shutdown(context.Background())
	}()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
