package Routes

import (
	"log"
	"net/http"

	"github.com/foolin/goview"
	"github.com/lfardell1/Go-Web-App-Blog/Database"
	models "github.com/lfardell1/Go-Web-App-Blog/Models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET / route logic here

	// Retrieve the users from the database
	var blogs []models.Post
	// Insert into blogs var
	blogs, err := Database.RetrieveBlogPosts()
	if err != nil {
		log.Fatal(err)
	}
	goview.Render(w, http.StatusOK, "index", goview.M{
		"Blogs": blogs,
	})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /login route logic here
	goview.Render(w, http.StatusOK, "about", goview.M{
		"title": "Hello",
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /register route logic here
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /logout route logic here
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /create route logic here
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /update route logic here
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /delete route logic here
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /user route logic here
}

func UserPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /user/posts route logic here
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /user/posts/:id route logic here
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /posts route logic here
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET /posts/:id route logic here
}
