package client

import (
	"go.uber.org/fx"

	"github.com/abvdasker/blog/client/db"
)

var Module = fx.Provide(
	db.New,
)
