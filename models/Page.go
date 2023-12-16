package models

// Post is a struct that represents a blog post.
type Page struct {
	Posts      []Post
	PrevPage   uint
	NextPage   uint
	CurPage    uint
	TotalPages uint
	TotalPosts uint
	ID         uint
}
