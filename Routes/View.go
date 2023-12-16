package Routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/foolin/goview"
	"github.com/lfardell1/Go-Web-App-Blog/middleware"
	"github.com/lfardell1/Go-Web-App-Blog/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract and parse the page number from the URL query parameters
	pageStr := r.URL.Query().Get("page")
	page := 1 // Default to page 1 if not provided or invalid
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}

	// Fetch blog posts for the requested page using your middleware function
	posts, err := middleware.PaginateBlogResults(uint(page))
	if err != nil {
		log.Printf("Error retrieving posts: %v", err)
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}

	// Render the retrieved posts using your template engine
	// Render index template
	data := struct {
		Title string
		Name  string
		Posts []models.Post // Assuming Post is your data structure
	}{
		Title: "Go Web App",
		Name:  "Leon",
		Posts: posts.Posts, // Pass your posts here
	}

	goview.Render(w, http.StatusOK, "index", data)
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
