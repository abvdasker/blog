package dal

import (
	"context"
	"database/sql"

	"github.com/abvdasker/blog/model"
)

const (
	readByUsernameQuery = `SELECT id, username, password_hash, salt, is_admin, created_at, updated_at FROM users WHERE username = $1`
)

type Users interface {
	ReadByUsername(
		ctx context.Context,
		username string,
	) (*model.User, error)
}

type users struct {
	db *sql.DB
}

func NewUsers(database *sql.DB) Users {
	return &users{
		db: database,
	}
}

func (a *users) ReadByUsername(
	ctx context.Context,
	username string,
) (*model.User, error) {
	row := a.db.QueryRowContext(
		ctx,
		readByUsernameQuery,
		username,
	)

	user := new(model.User)
	err := row.Scan(&user.ID, &user.Username, &user.IsAdmin, &user.Salt, &user.IsAdmin, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
