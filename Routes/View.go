package Routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/foolin/goview"
	"github.com/gorilla/mux"
	"github.com/lfardell1/Go-Web-App-Blog/middleware"
	"github.com/lfardell1/Go-Web-App-Blog/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	param := mux.Vars(r)
	if param["page"] == "" || param["page"] == "0" {
		param["page"] = "1"
	}
	page, err := strconv.Atoi(param["page"])
	if page <= 0 {
		page = 1
	}

	if err != nil {
		log.Printf("Error retrieving page: %v", err)
		http.Error(w, "Error retrieving page", http.StatusInternalServerError)
		return
	}

	// Fetch blog posts for the requested page using your middleware function
	posts, err := middleware.PaginateBlogResults(uint(page))
	if err != nil {
		log.Printf("Error retrieving posts: %v", err)
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}

	// Render index template
	data := struct {
		Title string
		Name  string
		Posts []models.Post
		Page  models.Page // Assuming Post is your data structure
	}{
		Title: "Go Web App",
		Name:  "Leon",
		Posts: posts.Posts,
		Page:  posts,
	}

	err = goview.Render(w, http.StatusOK, "index", data)
	if err != nil {
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {

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
