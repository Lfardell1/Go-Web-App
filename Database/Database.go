package Database

import (
	"fmt"
	"log"

	models "github.com/lfardell1/Go-Web-App-Blog/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"golang.org/x/crypto/bcrypt"
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

// Create User

func CreateUser(NewUser models.User) (models.User, error) {
	UserCollection := conn.Collection("Users")
	results := UserCollection.Find()
	var users []models.User
	err := results.All(&users)
	if err != nil {
		return NewUser, err

	}
	// check if email exists
	for _, user := range users {
		if user.Email == NewUser.Email {
			return user, err
		}
	}

	// if no email found insert user
	// using hashing system for password

	HashPassword, err := HashPassword(NewUser.Password)
	if err != nil {
		fmt.Println(err)

	}

	NewUser.Password = HashPassword

	_, err = UserCollection.Insert(NewUser)
	if err != nil {

		log.Fatalln(err)
		return NewUser, err
	}

	return NewUser, nil
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
	// find the comments related to this post and the author that made it
	for i := 0; i < len(posts); i++ {
		// find the comments related to this post and the author that made it
		comments, err := RetrieveComments(posts[i].ID)
		if err != nil {
			return models.Page{}, err
		}
		posts[i].Comments = comments
		author, err := RetrieveUser(posts[i].AuthorID)
		if err != nil {
			return models.Page{}, err
		}
		posts[i].Author = author.Username
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
		CurPage:    page,
		PrevPage:   page - 1,
		NextPage:   page + 1,
	}

	log.Printf("data: %+v", totalPages)
	log.Printf("data: %+v", page)

	return paginatedData, nil
}

func RetrieveComments(id uint) ([]models.Comment, error) {
	CommentCollection := conn.Collection("Comments")
	resultscomments := CommentCollection.Find(db.Cond{"PostID": id})

	var comments []models.Comment
	err := resultscomments.All(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func RetrieveUser(id uint) (models.User, error) {
	UserCollection := conn.Collection("Users")
	results := UserCollection.Find(db.Cond{"id": id})
	var user models.User
	err := results.One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
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
func CheckIfEmailExists(email string) bool {
	UserCollection := conn.Collection("Users")
	results := UserCollection.Find()
	var users []models.User
	err := results.All(&users)
	if err != nil {
		return false
	}

	for _, user := range users {
		if user.Email == email {
			return true
		}
	}
	return false
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func CheckIfUsernameExists(username string) bool {
	UserCollection := conn.Collection("Users")
	results := UserCollection.Find()
	var users []models.User
	err := results.All(&users)
	if err != nil {
		return false
	}

	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}
