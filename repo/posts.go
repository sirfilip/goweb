package repo

import (
	"goweb/model"
	"time"

	"github.com/jmoiron/sqlx"
)

func PostRepository(conn *sqlx.DB) *PostRepo {
	return &PostRepo{conn: conn}
}

type PostRepo struct {
	conn *sqlx.DB
}

func (pr *PostRepo) Posts() ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	err := pr.conn.Select(&posts, "SELECT * FROM POSTS")
	if nil != err {
		return posts, err
	}
	return posts, nil
}

func (pr *PostRepo) Create(post *model.Post) (*model.Post, error) {
	publishedAt := time.Now()
	result, err := pr.conn.Exec("INSERT INTO posts (title, content, published_at) VALUES (?, ?, ?)", post.Title, post.Content, publishedAt)
	if nil != err {
		return post, err
	}
	lastInsertId, err := result.LastInsertId()
	if nil != err {
		return post, err
	}
	post.ID = lastInsertId
	post.PublishedAt = publishedAt
	return post, nil
}
