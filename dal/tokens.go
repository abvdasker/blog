package dal

import (
	"context"
	"database/sql"

	"github.com/abvdasker/blog/model"
)

const (
	createToken = `INSERT INTO tokens (uuid, token, user_uuid, created_at, expires_at) VALUES ($1, $2, $3, $4, $5)`
	readByToken = `SELECT uuid, token, user_uuid, created_at, expires_at FROM tokens WHERE token = $1`
)

type Tokens interface {
	Create(
		ctx context.Context,
		token *model.Token,
	) error
	ReadByToken(
		ctx context.Context,
		tokenStr string,
	) (*model.Token, error)
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
		token.UUID,
		token.Token,
		token.UserUUID,
		token.CreatedAt,
		token.ExpiresAt,
	)
	return err
}

func (a *tokens) ReadByToken(
	ctx context.Context,
	tokenStr string,
) (*model.Token, error) {
	row := a.db.QueryRowContext(ctx, readByToken, tokenStr)

	token := new(model.Token)

	err := row.Scan(
		&token.UUID,
		&token.Token,
		&token.UserUUID,
		&token.CreatedAt,
		&token.ExpiresAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return token, nil
}
