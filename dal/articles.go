package dal

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/abvdasker/blog/model"
	"github.com/lib/pq"
)

const (
	readByDateQuery = `
SELECT 
  uuid, 
  title, 
  url_slug, 
  html, 
  tags, 
  created_at, 
  updated_at 
FROM articles 
WHERE 
  created_at > $1 
AND 
  created_at < $2 
ORDER BY created_at DESC
LIMIT $3 
OFFSET $4`
	createArticle = `INSERT INTO articles (uuid, title, html, url_slug, tags, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	readByURLSlug = `SELECT uuid, title, url_slug, html, tags, created_at, updated_at FROM articles WHERE url_slug = $1`
	updateArticle = `UPDATE articles SET title = $2, html = $3, url_slug = $4, tags = $5, updated_at = $6 WHERE uuid = $1`
	deleteArticle = `DELETE FROM articles WHERE uuid = $1`
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
	ReadByURLSlug(
		ctx context.Context,
		urlSlug string,
	) (*model.Article, error)
	Update(
		ctx context.Context,
		article *model.Article,
	) error
	Delete(
		ctx context.Context,
		articleUUID string,
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
		err := rows.Scan(&uuid, &title, &urlString, &html, pq.Array(&tags), &createdAt, &updatedAt)
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
		pq.Array(article.Base.Tags),
		article.Base.CreatedAt,
		article.Base.UpdatedAt,
	)
	return err
}

func (a *articles) ReadByURLSlug(
	ctx context.Context,
	urlSlug string,
) (*model.Article, error) {
	row := a.db.QueryRowContext(
		ctx,
		readByURLSlug,
		urlSlug,
	)
	article := model.Article{}

	err := row.Scan(
		&article.Base.UUID,
		&article.Base.Title,
		&article.Base.URLSlug,
		&article.HTML,
		pq.Array(article.Base.Tags),
		&article.Base.CreatedAt,
		&article.Base.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (a *articles) Update(
	ctx context.Context,
	article *model.Article,
) error {
	_, err := a.db.ExecContext(
		ctx,
		updateArticle,
		article.Base.UUID,
		article.Base.Title,
		article.HTML,
		article.Base.URLSlug,
		pq.Array(article.Base.Tags),
		article.Base.UpdatedAt,
	)
	return err
}

func (a *articles) Delete(
	ctx context.Context,
	articleUUID string,
) error {
	_, err := a.db.ExecContext(
		ctx,
		deleteArticle,
		articleUUID,
	)
	return err
}
