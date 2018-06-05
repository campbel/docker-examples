package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Print("starting up")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
	time.Sleep(time.Second)
	fmt.Println("ready")
	time.Sleep(time.Second * 3)
	fmt.Println("uh oh... something went wro")
	os.Exit(1)
}
