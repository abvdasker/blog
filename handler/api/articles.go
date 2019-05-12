package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/abvdasker/blog/dal"
)

type Articles interface {
	GetArticles() httprouter.Handle
}

type articles struct {
	articlesDAL dal.Articles
}

func NewArticles(articlesDAL dal.Articles) Articles {
	return &articles{
		articlesDAL: articlesDAL,
	}
}

func (a *articles) GetArticles() httprouter.Handle {
	return a.Handle
}

func (a *articles) Handle(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	then := time.Time{}
	now := time.Now()
	ctx := context.Background()
	articles, err := a.articlesDAL.ReadByDate(
		ctx,
		then, now,
		1000, 0,
	)
	if err != nil {
		respondErr(responseWriter, "error reading articles from database")
	}

	data, err := json.Marshal(articles)
	if err != nil {
		respondErr(responseWriter, "error serializing article data")
		return
	}
	responseWriter.Write(data)
}
