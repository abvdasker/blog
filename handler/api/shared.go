package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseRequest(rawRequest *http.Request, request interface{}) error {
	data, err := ioutil.ReadAll(rawRequest.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, request); err != nil {
		return err
	}
	return nil
}

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
