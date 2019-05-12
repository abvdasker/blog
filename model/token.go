package model

import (
	"time"

	"github.com/abvdasker/blog/lib"
)

const (
	defaultTokenExpiration = (time.Hour * 24)
)

type Token struct {
	ID     int    `json:"-"`
	Token  string `json:"token"`
	UserID int    `json:"userID"`

	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewToken(userID int, username string, salt string) *Token {
	now := time.Now()
	expiresAt := now.Add(defaultTokenExpiration)
	return &Token{
		Token:     lib.GenerateToken(username, salt, expiresAt),
		UserID:    userID,
		CreatedAt: now,
		ExpiresAt: expiresAt,
	}
}
