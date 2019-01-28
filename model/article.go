package model

type Article struct {
	Base BaseArticle `json:"base"`
	HTML string `json:"html"`
}
