package model

import (
	"time"

	"github.com/abvdasker/blog/lib"
)

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`

	Salt         string `json:"-"`
	PasswordHash string `json"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username string, password string, isAdmin bool) *User {
	salt := lib.RandomSalt64()
	now := time.Now()
	return &User{
		Username: username,
		IsAdmin:  isAdmin,

		Salt:         salt,
		PasswordHash: lib.HashPassword64(username, salt, password),

		CreatedAt: now,
		UpdatedAt: now,
	}
}
