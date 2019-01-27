package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Articles interface {
	Articles() httprouter.Handle
}

type articles struct {
}

func NewArticles() Articles {
	return &articles{}
}

func (a *articles) Articles() httprouter.Handle {
	return a.Handle
}

func (a *articles) Handle(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responseWriter.Write([]byte("Hello World"))
}
