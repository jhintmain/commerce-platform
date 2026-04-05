package server

import (
	"commerce-platform/internal/config"
	"commerce-platform/internal/router"

	"go.uber.org/zap"
)

type Server struct {
	cfg    *config.Config
	logger *zap.Logger
}

func New(cfg *config.Config, logger *zap.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Run() error {
	r := router.New()
	return r.Run()
}
