package model

import (
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`

	Salt         string `json:"-"`
	PasswordHash string `json"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username string, password string, isAdmin boll) *User {
	salt := lib.RandomSalt64()
	now := time.Now()
	return &User{
		Username: username,
		IsAdmin: isAdmin,

		Salt: salt,
		PasswordHash: lib.HashPassword64(username, salt, password),

		CreatedAt: now,
		UpdatedAt: now,
	}
}
