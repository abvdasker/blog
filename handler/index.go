package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IndexFn httprouter.Handle

func NewIndex() IndexFn {
	return func(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		http.ServeFile(responseWriter, request, "static/html/index.html")
	}
}
