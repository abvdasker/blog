package server

import (
	"net/http"

	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/server/handler"
)

type Server interface {
	Start() error
}

type server struct {
	cfg        *config.Config
	base       *http.Server
	apiHandler handler.APIHandler
}

func New(
	cfg *config.Config,
	apiHandler handler.APIHandler,
	middlewareHandler handler.MiddlewareHandler,
) Server {
	return &server{
		cfg: cfg,
		base: &http.Server{
			Addr:        cfg.Server.Hostport,
			ReadTimeout: cfg.Server.Timeout,
			Handler: middlewareHandler,
		},
		apiHandler: apiHandler,
	}
}

func (s *server) Start() error {
	staticFilesDir := http.Dir("static")
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(staticFilesDir)))
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		http.ServeFile(responseWriter, request, "static/html/index.html")
	})
	http.Handle("/api", s.apiHandler)
	return s.base.ListenAndServe()
}

func Start(server Server) error {
	return server.Start()
}
