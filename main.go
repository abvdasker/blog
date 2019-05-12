package main

import (
	"go.uber.org/fx"

	"github.com/abvdasker/blog/client"
	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/dal"
	"github.com/abvdasker/blog/handler"
	"github.com/abvdasker/blog/handler/api"
	"github.com/abvdasker/blog/lib"
	"github.com/abvdasker/blog/server"
)

func main() {
	app := fx.New(
		handler.Module,
		api.Module,
		client.Module,
		dal.Module,
		fx.Provide(
			config.Load,
			server.NewRouter,
			server.New,
			lib.NewLogger,
		),
		fx.Invoke(server.Start),
	)

	app.Run()
}
