package main

import (
	"go.uber.org/fx"

	"github.com/abvdasker/blog/client"
	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/dal"
	"github.com/abvdasker/blog/handler"
	"github.com/abvdasker/blog/handler/api"
	"github.com/abvdasker/blog/handler/api/middleware"
	"github.com/abvdasker/blog/lib"
	"github.com/abvdasker/blog/server"
)

func main() {
	app := fx.New(
		handler.Module,
		api.Module,
		client.Module,
		dal.Module,
		server.Module,
		fx.Provide(
			middleware.NewAuth,
			config.Load,
			lib.NewLogger,
		),
		fx.Invoke(server.Start),
	)

	app.Run()
}
