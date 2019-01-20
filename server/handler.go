package server

import (
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("Hello World"))
}
