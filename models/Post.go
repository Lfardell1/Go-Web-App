package models

import "time"

type Post struct {
	ID          uint      `db:"id,omitempty"`
	DateCreated time.Time `db:"Created_at"`
	Title       string    `db:"Title"`
	ImageURL    string    `db:"Photo"`
	Author      string    `db:"Summary"`
	AuthorID    uint      `db:"AuthorID"`
}
