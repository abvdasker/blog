package main

import (
	"go.uber.org/fx"

	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/server"
	"github.com/abvdasker/blog/server/handler"
	"github.com/abvdasker/blog/lib"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.Load,
			handler.NewAPIHandler,
			handler.NewMiddlewareHandler,
			server.New,
			lib.NewLogger,
		),
		fx.Invoke(server.Start),
	)

	app.Run()
}
