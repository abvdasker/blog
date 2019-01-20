package lib

import (
	"go.uber.org/zap"

	"github.com/abvdasker/blog/config"
)

func NewLogger(cfg *config.Config) (*zap.SugaredLogger, error) {
	logger, err := cfg.Logger.Build()
	if err != nil {
		return nil, err
	}
	logger.Sugar().With(zap.String("test", "hello world")).Info("TEST MESSAGE")
	return logger.Sugar(), nil
}
