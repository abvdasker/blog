package dal

import (
	"context"

	"database/sql"

	"github.com/abvdasker/blog/client/db"
)

const (
	readByDateQuery = `SELECT * FROM articles WHERE CREATED_AT > ? AND CREATED_AT < ? LIMIT ? OFFSET ?`
)

type Articles interface {
	ReadByDate(
		ctx context.Context,
		start, end time.Time,
		limit, offset int,
	) ([]*model.Article, error)
}

type articles struct {
	readByDate *db.
}

func NewArticles(database *sql.DB) Articles {
	readByDate := database.Prepare(
		readByDateQuery,
	)

	return &articles{
		readByDate: readByDate,
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

	rows, err := a.readByDate.QueryContext(ctx, start, end, limit, offset)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	
}
