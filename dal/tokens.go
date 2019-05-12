package dal

import (
	"context"
	"database/sql"

	"github.com/abvdasker/blog/model"
)

const (
	createToken = `INSERT INTO tokens (token, user_id, created_at, expires_at) VALUES (?, ?, ?, ?)`
)

type Tokens interface {
	Create(
		ctx context.Context,
		token *model.Token,
	) error
}

type tokens struct {
	db *sql.DB
}

func NewTokens(database *sql.DB) Tokens {
	return &tokens{
		db: database,
	}
}

func (a *tokens) Create(
	ctx context.Context,
	token *model.Token,
) error {
	_, err := a.db.ExecContext(
		ctx,
		createToken,
		token.Token,
		token.UserID,
		token.CreatedAt,
		token.ExpiresAt,
	)
	return err
}
