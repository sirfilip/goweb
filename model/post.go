package model

import "time"

type Post struct {
	ID          int64     `db:"id"`
	Title       string    `db:"title"`
	Content     string    `db:"content"`
	PublishedAt time.Time `db:"published_at"`
}
