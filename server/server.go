package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/handler"
)

type Server interface {
	Start() error
}

type server struct {
	cfg    *config.Config
	base   *http.Server
	logger *zap.SugaredLogger
}

func New(cfg *config.Config, router *httprouter.Router, middleware handler.Middleware, logger *zap.SugaredLogger) Server {
	return &server{
		cfg:    cfg,
		logger: logger,
		base: &http.Server{
			Addr:        cfg.Server.Hostport,
			ReadTimeout: cfg.Server.Timeout,
			Handler:     middleware.Wrap(router),
		},
	}
}

func (s *server) Start() error {
	s.logger.With(zap.String("hostport", s.cfg.Server.Hostport)).Info("Server starting")
	return s.base.ListenAndServe()
}

func Start(server Server) error {
	return server.Start()
}
