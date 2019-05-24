package lib

import (
	"testing"

	"go.uber.org/zap"

	"github.com/abvdasker/blog/config"
)

func TestNewLogger(t *testing.T) {
	cfg := &config.Config{
		Logger: zap.Config{
			Encoding: "console",
		},
	}
	logger, err := NewLogger(cfg)

	if err != nil {
		t.Fatalf("building the logger returned an error: %s", err.Error())
	}
	if logger == nil {
		t.Fatalf("the logger was nil but no error was thrown")
	}
}

func TestNewLoggerError(t *testing.T) {
	cfg := &config.Config{}
	logger, err := NewLogger(cfg)

	if err == nil {
		t.Fatalf("building the logger should return an error")
	}
	if logger != nil {
		t.Fatalf("an error was thrown but the logger was not nil")
	}
}
