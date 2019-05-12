package dal

import (
	"context"
	"database/sql"

	"github.com/abvdasker/blog/model"
)

const (
	readByUsernameQuery = `SELECT uuid, username, password_hash, salt, is_admin, created_at, updated_at FROM users WHERE username = $1`

	createUser = `INSERT INTO users (username, password_hash, salt, is_admin, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
)

type Users interface {
	ReadByUsername(
		ctx context.Context,
		username string,
	) (*model.User, error)
	Create(
		ctx context.Context,
		user *model.User,
	) error
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
	err := row.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.Salt, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (a *users) Create(ctx context.Context, user *model.User) error {
	_, err := a.db.ExecContext(
		ctx,
		createUser,
		user.Username,
		user.PasswordHash,
		user.Salt,
		user.IsAdmin,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}
