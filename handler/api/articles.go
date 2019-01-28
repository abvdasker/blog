package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/abvdasker/blog/model"
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
	articles := []model.Article{
		{
			Base: model.BaseArticle{
				ID:        1,
				Title:     "The bandwagon",
				URLSlug:   "the-bandwagon",
				Tags:      []string{"introduction", "welcome"},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			HTML: `<p>Welcome to my blog.</p>
<p>Today I started my blog and it was good. Read it and enjoy the blog. Everything on here is good. Rest assured. No bad quality.</p>
`,
		},
		{
			Base: model.BaseArticle{
				ID:        2,
				Title:     "Another article",
				URLSlug:   "the-bandwagon",
				Tags:      []string{"introduction", "welcome"},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			HTML: `<p>Welcome to my blog.</p>
<p>Today I started my blog and it was good. Read it and enjoy the blog. Everything on here is good. Rest assured. No bad quality.</p>
`,
		},
	}
	data, err := json.Marshal(articles)
	if err != nil {
		respondErr(responseWriter, "error serializing article data")
		return
	}
	responseWriter.Write(data)
}

func respondErr(responseWriter http.ResponseWriter, msg string) {
	responseWriter.WriteHeader(500)
	responseWriter.Write([]byte(msg))
}
