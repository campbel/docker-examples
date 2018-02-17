package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	fmt.Println("Waiting for docker stop signal")
	<-sigs
	fmt.Println("Signal received!")
}
