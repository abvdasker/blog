package main

import (
	"go.uber.org/fx"

	"github.com/abvdasker/blog/config"
	"github.com/abvdasker/blog/server"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.Load,
			server.New,
		),
		fx.Invoke(server.Start),
	)

	app.Run()
}
