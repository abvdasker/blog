package model

import (
	"time"
)

type BaseArticle struct {
	UUID    string   `json:"uuid"`
	Title   string   `json:"title"`
	URLSlug string   `json:"urlSlug"`
	Tags    []string `json:"tags"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
