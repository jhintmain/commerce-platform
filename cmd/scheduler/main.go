package main

import (
	"commerce-platform/internal/common/logpkg"
	"time"
)

func main() {

	//cfg := config.Load()

	logger := logpkg.New("scheduler")
	logger.Info("Starting scheduler")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		logger.Info("Starting thick")
	}
}
