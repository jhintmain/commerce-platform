package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Starting scheduler")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Starting thick")
	}
}
