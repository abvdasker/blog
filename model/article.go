package model

import (
	"strings"
	"time"

	"github.com/abvdasker/blog/lib/uuid"
)

var articles = map[string]bool{
	"a":   true,
	"the": true,
}

type BaseArticle struct {
	UUID    string   `json:"uuid"`
	Title   string   `json:"title"`
	URLSlug string   `json:"urlSlug"`
	Tags    []string `json:"tags"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Article struct {
	Base BaseArticle `json:"base"`
	HTML string      `json:"html"`
}

func NewArticleWithUUID(uuid, title, html string, tags []string) *Article {
	now := time.Now()
	return &Article{
		Base: BaseArticle{
			UUID:    uuid,
			Title:   title,
			URLSlug: generateURLSlug(title),
			Tags:    tags,

			CreatedAt: now,
			UpdatedAt: now,
		},
		HTML: html,
	}
}

func NewArticle(title string, html string, tags []string) *Article {
	return NewArticleWithUUID(
		uuid.New().String(),
		title,
		html,
		tags,
	)
}

func generateURLSlug(title string) string {
	lowered := strings.ToLower(title)
	exploded := strings.Split(lowered, " ")
	explodedNoArticles := filterArticles(exploded)
	return strings.Join(explodedNoArticles, "-")
}

func filterArticles(words []string) []string {
	filtered := make([]string, 0, len(words))
	for _, word := range words {
		if _, ok := articles[word]; ok {
			continue
		}
		filtered = append(filtered, word)
	}
	return filtered
}
