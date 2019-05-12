package api

import (
	"net/http"
)

func respondNotFound(responseWriter http.ResponseWriter, msg string) {
	http.Error(responseWriter, msg, http.StatusNotFound)
}

func respondErr(responseWriter http.ResponseWriter, msg string) {
	http.Error(responseWriter, msg, http.StatusInternalServerError)
}

func respondUnauthorized(responseWriter http.ResponseWriter, msg string) {
	http.Error(responseWriter, msg, http.StatusUnauthorized)
}

func respondStatusMsg(responseWriter http.ResponseWriter, status int, msg string) {
	responseWriter.WriteHeader(status)
	responseWriter.Write([]byte(msg))
}
