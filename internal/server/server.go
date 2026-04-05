package server

import (
	"commerce-platform/internal/config"
	"commerce-platform/internal/db/generate/db"
	"commerce-platform/internal/router"

	"go.uber.org/zap"
)

type Server struct {
	cfg    *config.Config
	logger *zap.Logger
	query  *db.Query
}

func New(cfg *config.Config, logger *zap.Logger, query *db.Query) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		query:  query,
	}
}

func (s *Server) Run() error {
	r := router.New()
	return r.Run()
}
