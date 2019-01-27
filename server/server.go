package server

import (
	"net/http"

	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/handler"
	"github.com/julienschmidt/httprouter"
)

type Server interface {
	Start() error
}

type server struct {
	cfg  *config.Config
	base *http.Server
}

func New(cfg *config.Config, router *httprouter.Router, middleware handler.Middleware) Server {
	return &server{
		cfg: cfg,
		base: &http.Server{
			Addr:        cfg.Server.Hostport,
			ReadTimeout: cfg.Server.Timeout,
			Handler:     middleware.Wrap(router),
		},
	}
}

func (s *server) Start() error {
	return s.base.ListenAndServe()
}

func Start(server Server) error {
	return server.Start()
}
