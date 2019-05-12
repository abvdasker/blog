package server

import (
	"github.com/julienschmidt/httprouter"

	"github.com/abvdasker/blog/handler"
	"github.com/abvdasker/blog/handler/api"
)

func NewRouter(
	staticHandler handler.Static,
	articlesHandler api.Articles,
	usersHandler api.Users,
) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/articles", articlesHandler.GetArticles())
	router.POST("/api/users/login", usersHandler.Login())
	router.POST("/api/users", usersHandler.Create())
	router.GET("/static/*filepath", staticHandler.Static())
	router.GET("/", staticHandler.Index())

	return router
}
