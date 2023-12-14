package Database

import (
	"log"

	models "github.com/lfardell1/Go-Web-App-Blog/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

type ConnectionURL struct {
	User     string
	Password string
	Host     string
	Database string
	Options  map[string]string
}

var settings = mysql.ConnectionURL{
	Database: "GoWebApp",
	Password: "go",
	Host:     "127.0.0.1",
	User:     "go",
}

var conn db.Session

func init() {
	var err error
	conn, err = mysql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}

}

func RetrieveUsers() ([]models.User, error) {
	UserCollection := conn.Collection("Users")
	results := UserCollection.Find()

	var users []models.User
	err := results.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func RetrieveBlogPost() ([]models.Post, error) {
	BlogCollection := conn.Collection("Blogs")
	resultsblogs := BlogCollection.Find()

	var posts []models.Post
	err := resultsblogs.All(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPaginatedResults(page uint) ([]models.Post, error) {

	// establishing variables
	var blog_posts []models.Post
	return blog_posts, nil
}
