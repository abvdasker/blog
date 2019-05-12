package api

import (
	"net/http"
)

func respondNotFound(responseWriter http.ResponseWriter, msg string) {
	respondStatusMsg(responseWriter, 404, msg)
}

func respondErr(responseWriter http.ResponseWriter, msg string) {
	respondStatusMsg(responseWriter, 500, msg)
}

func respondUnauthorized(responseWriter http.ResponseWriter, msg string) {
	respondStatusMsg(responseWriter, 401, msg)
}

func respondStatusMsg(responseWriter http.ResponseWriter, status int, msg string) {
	responseWriter.WriteHeader(status)
	responseWriter.Write([]byte(msg))
}
