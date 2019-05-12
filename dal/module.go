package dal

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewArticles,
	NewUsers,
	NewTokens,
)
