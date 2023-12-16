package models

import "time"

type Comment struct {
	ID          uint      `db:"id,omitempty"`
	PostID      uint      `db:"PostID"`
	AuthorID    uint      `db:"UserID"`
	DateCreated time.Time `db:"Created_at"`
	Content     string    `db:"Content"`
	Author      string
}
