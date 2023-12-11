package models

type User struct {
	user_id  int    `db:"user_id,omitempty"`
	Username string `db:"username"`
	Password string `db:"password"`
	Name     string `db:"name"`
	Email    string `db:"email"`
}
