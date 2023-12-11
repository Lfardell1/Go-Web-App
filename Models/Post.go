package models

type Post struct {
	ID       int    `db:"id,omitempty"`
	UserId   int    `db:"UserId"`
	Title    string `db:"title"`
	ImageURL string `db:"image_url"`
	Content  string `db:"content"`
	Author   string `db:"author"`
}
