package database

import "time"

type ArticleData struct {
	Title       string
	AuthorID    int64
	AuthorName  string
	PostTime    time.Time
	UpdatedTime time.Time
	Status      string
}
