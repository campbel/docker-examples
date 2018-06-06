package main

import (
	"fmt"
	"net/http"
	"os"
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
	done := make(chan bool)
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGTERM)
	// go func() {
	// 	<-sigs // this blocks until SIGTERM is sent
	// 	fmt.Println("shutting down...")
	// 	server.Shutdown(context.Background())
	// 	done <- true
	// }()

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		os.Exit(1)
	}

	<-done
	fmt.Println("shutdown")
}
