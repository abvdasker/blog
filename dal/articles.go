package dal

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/abvdasker/blog/model"
)

const (
	readByDateQuery = `SELECT uuid, title, url_string, html, tags, created_at, updated_at FROM articles WHERE CREATED_AT > $1 AND CREATED_AT < $2 LIMIT $3 OFFSET $4`
	createArticle = `INSERT INTO articles (uuid, title, html, url_slug, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
)

type Articles interface {
	ReadByDate(
		ctx context.Context,
		start, end time.Time,
		limit, offset int,
	) ([]*model.Article, error)
	Create(
		ctx context.Context,
		article *model.Article,
	) error
}

type articles struct {
	db *sql.DB
}

func NewArticles(database *sql.DB) Articles {
	return &articles{
		db: database,
	}
}

func (a *articles) ReadByDate(
	ctx context.Context,
	start, end time.Time,
	limit, offset int,
) ([]*model.Article, error) {
	if !start.Before(end) {
		return nil, errors.New("start time must be earlier than end time")
	}

	rows, err := a.db.QueryContext(ctx, readByDateQuery, start, end, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*model.Article, 0, 0)

	for rows.Next() {
		var (
			uuid      string
			title     string
			urlString string
			html      string
			tags      []string
			createdAt time.Time
			updatedAt time.Time
		)
		err := rows.Scan(&uuid, &title, &urlString, &html, &tags, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		baseArticle := model.BaseArticle{
			UUID:      uuid,
			Title:     title,
			URLSlug:   urlString,
			Tags:      tags,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		article := &model.Article{
			Base: baseArticle,
			HTML: html,
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (a *articles) Create(
	ctx context.Context,
	article *model.Article,
) error {
	_, err := a.db.ExecContext(
		ctx,
		createArticle,
		article.Base.UUID,
		article.Base.Title,
		article.HTML,
		article.Base.URLSlug,
		article.Base.Tags,
		article.Base.CreatedAt,
		article.Base.UpdatedAt,
	)
	return err
}
