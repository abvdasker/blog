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

type Client interface {
	Close() error
}

type client struct {
	db *sql.DB
}

func New(cfg *config.Config) (Client, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(databaseURLTemplate, cfg.DB.Hostport),
	)

	return &client{
		db: db,
	}, err
}

func (c *client) Close() error {
	return c.db.Close()
}
