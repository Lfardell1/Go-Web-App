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

// Establish a connection to the database
var conn db.Session

func init() {
	var err error
	conn, err = mysql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}

}

// Retrieve all users
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

// Single blog
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

func PaginateBlogResults(page uint, postsPerPage uint) (models.Page, error) {
	blogsTable := conn.Collection("Blogs")

	// Query to retrieve paginated blog posts
	var posts []models.Post

	// Calculate offset based on the requested page
	offset := (page - 1) * postsPerPage

	// Fetch paginated data
	err := blogsTable.Find().OrderBy("Created_at DESC").Limit(int(postsPerPage)).Offset(int(offset)).All(&posts)
	if err != nil {
		return models.Page{}, err
	}

	// Count total number of posts
	totalPostsCount, err := blogsTable.Find().Count()
	if err != nil {
		return models.Page{}, err
	}

	// Calculate total pages, considering remaining posts
	totalPages := totalPostsCount / uint64(postsPerPage)
	if totalPostsCount%uint64(postsPerPage) != 0 {
		totalPages++
	}

	// Create a paginated response
	paginatedData := models.Page{
		Posts:      posts,
		TotalPosts: uint(totalPostsCount),
		TotalPages: uint(totalPages),
	}

	return paginatedData, nil
}

func RetrievePost(id int) ([]models.Post, error) {
	BlogCollection := conn.Collection("Blogs")
	resultsblogs := BlogCollection.Find(db.Cond{"id": id})

	var post []models.Post
	err := resultsblogs.One(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}
