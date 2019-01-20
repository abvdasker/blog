package server

import (
	"net/http"

	"github.com/abvdasker/blog/config"
)

type Server interface {
	Start() error
}

type server struct {
	cfg  *config.Config
	base *http.Server
}

func New(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
		base: &http.Server{
			Addr:        cfg.Server.Hostport,
			ReadTimeout: cfg.Server.Timeout,
			Handler:     new(handler),
		},
	}
}

func (s *server) Start() error {
	return s.base.ListenAndServe()
}

func Start(server Server) error {
	return server.Start()
}
