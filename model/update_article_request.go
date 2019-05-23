package model

type UpdateArticleRequest struct {
	UUID  string   `json:"uuid"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	HTML  string   `json:"html"`
}

func (r *UpdateArticleRequest) ToArticle() *Article {
	return NewArticleWithUUID(
		r.UUID,
		r.Title,
		r.HTML,
		r.Tags,
	)
}
