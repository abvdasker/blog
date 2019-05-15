package model

import (
	"time"

	"github.com/abvdasker/blog/lib"
	"github.com/abvdasker/blog/lib/uuid"
)

const (
	defaultTokenExpiration = (time.Hour * 24)
)

type Token struct {
	UUID     string `json:"-"`
	Token    string `json:"token"`
	UserUUID string `json:"userUUID"`

	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewToken(userUUID, username, salt string) *Token {
	now := time.Now()
	expiresAt := now.Add(defaultTokenExpiration)
	return &Token{
		UUID:      uuid.New().String(),
		Token:     lib.GenerateToken(username, salt, expiresAt),
		UserUUID:  userUUID,
		CreatedAt: now,
		ExpiresAt: expiresAt,
	}
}
