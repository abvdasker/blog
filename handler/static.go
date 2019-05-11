package handler

import (
	"net/http"

	"go.uber.org/zap"
	"github.com/julienschmidt/httprouter"
)

type Static interface {
	Static() httprouter.Handle
	Index() httprouter.Handle
}

type static struct {
	logger *zap.SugaredLogger
}

func NewStatic(logger *zap.SugaredLogger) Static {
	return &static{
		logger: logger,
	}
}

func (s *static) Static() httprouter.Handle {
	staticFilesDir := http.Dir("static")
	return func(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		s.logger.With(zap.String("path", request.URL.Path)).Info("static request")
		http.StripPrefix("/static/", http.FileServer(staticFilesDir)).ServeHTTP(responseWriter, request)
	}
}

func (s *static) Index() httprouter.Handle {
	return func(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		s.logger.With(zap.String("path", request.URL.Path)).Info("index request")
		http.ServeFile(responseWriter, request, "static/html/index.html")
	}
}
