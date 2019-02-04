package db

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/abvdasker/blog/config"
)

const (
	databaseURLTemplate = "postgresql://blog@%s/blog"
)

func New(cfg *config.Config) (*sql.DB, error) {
	return sql.Open(
		"postgres",
		fmt.Sprintf(databaseURLTemplate, cfg.DB.Hostport),
	)
}
