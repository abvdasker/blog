package db

import (
	"fmt"
	"database/sql"

	"go.uber.org/zap"
	_ "github.com/lib/pq"

	"github.com/abvdasker/blog/config"
)

const (
	databaseURLTemplate = "postgresql://blog@%s/blog?sslmode=%s"
)

func New(cfg *config.Config, logger *zap.SugaredLogger) (*sql.DB, error) {
	logger.Info("opening database connection")
	return sql.Open(
		"postgres",
		fmt.Sprintf(databaseURLTemplate, cfg.DB.Hostport, cfg.DB.DisableSSLStr()),
	)
}
