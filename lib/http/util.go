package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseRequest(rawRequest *http.Request, request interface{}) error {
	data, err := ioutil.ReadAll(rawRequest.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, request); err != nil {
		return err
	}
	return nil
}

func RespondBadRequest(responseWriter http.ResponseWriter, msg string) {
	http.Error(responseWriter, msg, http.StatusBadRequest)
}

func RespondNotFound(responseWriter http.ResponseWriter, msg string) {
	http.Error(responseWriter, msg, http.StatusNotFound)
}

func RespondErr(responseWriter http.ResponseWriter, msg string) {
	http.Error(responseWriter, msg, http.StatusInternalServerError)
}

func RespondUnauthorized(responseWriter http.ResponseWriter, msg string) {
	http.Error(responseWriter, msg, http.StatusUnauthorized)
}
