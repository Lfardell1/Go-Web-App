package models

type User struct {
	ID         uint   `db:"id,omitempty"`
	Username   string `db:"Username"`
	Password   string `db:"Password"`
	Name       string `db:"Name"`
	Email      string `db:"Email"`
	ProfilePic string `db:"ProfilePhoto"`
	About      string `db:"About"`
}
