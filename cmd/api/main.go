package main

import (
	"commerce-platform/internal/common/logpkg"
	"commerce-platform/internal/config"
	"commerce-platform/internal/server"
	"log"
)

func main() {
	cfg := config.Load()

	logger := logpkg.New("api")
	srv := server.New(cfg, logger)

	log.Println("[API] Starting server", cfg.AppPort)

	if err := srv.Run(); err != nil {
		log.Fatalf("[API] server run error : %v", err)
	}

}
