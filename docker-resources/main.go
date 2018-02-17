package main

import "time"

func main() {
	var bigOleArray [][]string

	ticker := time.NewTicker(time.Second)

	for _ = range ticker.C {
		bigOleArray = append(bigOleArray, make([]string, 65536))
	}
}
