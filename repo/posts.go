package repo

import (
	"goblog/model"
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

func (pr *PostRepo) Create(title, content string) (*model.Post, error) {
	post := &model.Post{Title: title, Content: content}
	result, err := pr.conn.Exec("INSERT INTO posts (title, content, published_at) VALUES (?, ?, ?)", title, content, time.Now())
	if nil != err {
		return post, err
	}
	lastInsertId, err := result.LastInsertId()
	if nil != err {
		return post, err
	}
	post.ID = lastInsertId
	return post, nil
}
