package handler

import (
	"net/http"
)

type APIHandler interface {
	http.Handler
}

type apiHandler struct {
}

func NewAPIHandler() APIHandler {
	return &apiHandler{}
}

func (h *apiHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("Hello World"))
}
