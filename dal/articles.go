package dal

import (
	"context"

	"github.com/abvdasker/blog/client/db"
)

type Articles interface {
	ReadByDate(
		ctx context.Context,
		start, end time.Time,
		limit, offset int,
	) ([]*model.Article, error)
}

type articles struct {
	database db.Client
}

func NewArticles(database db.Client) Articles {
	return &articles{
		database: database,
	}
}

func (a *articles) ReadByDate(
	ctx context.Context,
	start, end time.Time,
	limit, offset int,
) ([]*model.Article, error) {
	if !start.Before(end) {
		return nil, erorrs.New("start time must be earlier than end time")
	}
	
}
