package server

import (
	"github.com/julienschmidt/httprouter"

	"github.com/abvdasker/blog/handler"
	"github.com/abvdasker/blog/handler/api"
)

func NewRouter(
	staticHandler handler.Static,
	articlesHandler api.Articles,
) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/articles", articlesHandler.Articles())
	router.GET("/static/*filepath", staticHandler.Static())
	router.GET("/", staticHandler.Index())

	return router

}
