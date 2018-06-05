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
	fmt.Println("starting up...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	server := http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	// Listen for SIGTERM and shutdown the server.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		<-sigs // this blocks until SIGTERM is sent
		fmt.Println("shutting down...")
		server.Shutdown(context.Background())
	}()

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
