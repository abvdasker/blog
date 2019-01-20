package server

import (
	"net/http"

	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/server/handler"
)

const (
	staticFilesPath = "static/"
)

type Server interface {
	Start() error
}

type server struct {
	cfg  *config.Config
	base *http.Server
	apiHandler handler.APIHandler
}

func New(cfg *config.Config, apiHandler handler.APIHandler) Server {
	return &server{
		cfg: cfg,
		base: &http.Server{
			Addr:        cfg.Server.Hostport,
			ReadTimeout: cfg.Server.Timeout,
		},
		apiHandler: apiHandler,
	}
}

func (s *server) Start() error {
	staticFilesDir := http.Dir("static")
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(staticFilesDir)))
	http.HandleFunc("/index.html", func(responseWriter http.ResponseWriter, request *http.Request) {
		http.ServeFile(responseWriter, request, "static/html/index.html")
	})
	http.Handle("/api", s.apiHandler)
	return s.base.ListenAndServe(s.cfg.Server.Hostport, nil)
}

func Start(server Server) error {
	return server.Start()
}
