package model

type CreateArticleRequest struct {
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	HTML  string   `json:"html"`
}

func (r *CreateArticleRequest) ToArticle() *Article {
	return NewArticle(
		r.Title,
		r.HTML,
		r.Tags,
	)
}
