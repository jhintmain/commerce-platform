package main

import (
	"commerce-platform/internal/common/logpkg"
	"commerce-platform/internal/config"
	"commerce-platform/internal/db/generate/db"
	mysqlpgk "commerce-platform/internal/infrastructure/mysql"
	"commerce-platform/internal/server"
	"log"

	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	logger := logpkg.New("api")

	mysqlDB, err := mysqlpgk.New(cfg.MySQLDSN)
	if err != nil {
		logger.Error("mysql connect failed", zap.Error(err))
	}
	defer mysqlDB.Close()

	query := db.New(mysqlDB)
	srv := server.New(cfg, logger, query)

	log.Println("[API] Starting server", cfg.AppPort)

	if err := srv.Run(); err != nil {
		log.Fatalf("[API] server run error : %v", err)
	}

}
