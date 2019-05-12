package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Static interface {
	Static() httprouter.Handle
	Index() httprouter.Handle
	CMS() httprouter.Handle
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

func (s *static) CMS() httprouter.Handle {
	return func(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		path := request.URL.Path
		s.logger.With(
			zap.String("path", request.URL.Path),
		).Info("cms request")
		filepath := fmt.Sprintf("static/html/%s", path)
		http.ServeFile(responseWriter, request, filepath)
	}
}
