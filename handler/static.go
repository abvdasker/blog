package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Static interface {
	Static() httprouter.Handle
	Index() httprouter.Handle
}

type static struct {
}

func NewStatic() Static {
	return &static{}
}

func (s *static) Static() httprouter.Handle {
	staticFilesDir := http.Dir("static")
	return func(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		http.StripPrefix("/static/", http.FileServer(staticFilesDir)).ServeHTTP(responseWriter, request)
	}
}

func (s *static) Index() httprouter.Handle {
	return func(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		http.ServeFile(responseWriter, request, "static/html/index.html")
	}
}
