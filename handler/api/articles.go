package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"github.com/abvdasker/blog/dal"
	"github.com/abvdasker/blog/model"
)

type Articles interface {
	GetArticles() httprouter.Handle
	CreateArticle() httprouter.Handle
}

type articles struct {
	articlesDAL dal.Articles
	logger      *zap.SugaredLogger
}

func NewArticles(articlesDAL dal.Articles, logger *zap.SugaredLogger) Articles {
	return &articles{
		articlesDAL: articlesDAL,
		logger:      logger,
	}
}

func (a *articles) GetArticles() httprouter.Handle {
	return a.HandleGetArticles
}

func (a *articles) CreateArticle() httprouter.Handle {
	return a.HandleCreateArticle
}

func (a *articles) HandleGetArticles(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	then := time.Time{}
	now := time.Now()
	ctx := context.Background()
	articles, err := a.articlesDAL.ReadByDate(
		ctx,
		then, now,
		1000, 0,
	)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("error reading articles to database")
		respondErr(responseWriter, "error reading articles from database")
		return
	}

	data, err := json.Marshal(articles)
	if err != nil {
		respondErr(responseWriter, "error serializing article data")
		return
	}
	responseWriter.Write(data)
}

func (a *articles) HandleCreateArticle(responseWriter http.ResponseWriter, rawRequest *http.Request, params httprouter.Params) {
	ctx := context.Background()
	request, err := parseCreateArticleRequest(rawRequest)
	if err != nil {
		respondErr(responseWriter, "failed to parse request")
		return
	}
	article := request.ToArticle()

	err = a.articlesDAL.Create(
		ctx,
		article,
	)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("error writing article to database")
		respondErr(responseWriter, "error writing article to database")
		return
	}

	data, err := json.Marshal(article)
	if err != nil {
		respondErr(responseWriter, "error serializing article data")
		return
	}
	responseWriter.Write(data)
}

func parseCreateArticleRequest(rawRequest *http.Request) (*model.CreateArticleRequest, error) {
	request := new(model.CreateArticleRequest)

	if err := parseRequest(rawRequest, request); err != nil {
		return nil, err
	}
	return request, nil
}
