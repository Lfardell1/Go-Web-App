package Routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	Helpers "github.com/lfardell1/Go-Web-App-Blog/Helpers"
	"github.com/lfardell1/Go-Web-App-Blog/middleware"
	"github.com/lfardell1/Go-Web-App-Blog/models"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RequestHandler struct{}

func GetTime(w http.ResponseWriter, r *http.Request) {

}

func ReturnLoginForm(w http.ResponseWriter, r *http.Request) {

}

func ReturnSignupForm(w http.ResponseWriter, r *http.Request) {

}

func RenderBlogs(w http.ResponseWriter, r *http.Request) {

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

	Helpers.RenderBlogPostsPage(w, page)

}

func Right(w http.ResponseWriter, r *http.Request) {
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

	// Render the retrieved posts using your template engine
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

	// Parse the HTML template
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the updated data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func PaginationHandler(w http.ResponseWriter, r *http.Request) {

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch users data from a database or provide dummy data

}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch posts data from a database or provide dummy data

}

// Implement other handler functions similarly

func GetPost(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch post data from a database or provide dummy data

}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch user posts data from a database or provide dummy data

}

func GetUserPost(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch a specific user post data from a database or provide dummy data

}

// LoginAPIHandler handles the login functionality
func Login(w http.ResponseWriter, r *http.Request) {
	// Logic for handling user login, authentication, and generating tokens
	// Implement your login logic here, handle user authentication, and generate tokens if successful
}

// Logout handles user logout
func Logout(w http.ResponseWriter, r *http.Request) {
	// Logic for handling user logout and invalidating tokens
	// Implement your logout logic here, invalidate user tokens, and perform any necessary actions
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Logic for handling user logout and invalidating tokens
	// Implement your logout logic here, invalidate user tokens, and perform any necessary actions
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Logic for handling user logout and invalidating tokens
	// Implement your logout logic here, invalidate user tokens, and perform any necessary actions
}

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	// Logic for user registration, validating input data, and storing user information
	// Implement your registration logic here, validate user input, store user data, and return appropriate responses
}

// DeletePost handles deleting a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Logic for deleting a post by ID
	// Implement your logic to delete a post by ID from a database or any source
}

// UpdatePost handles updating a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Logic for updating a post by ID
	// Implement your logic to update a post by ID with the data from the request body
}

// UpdateUser handles updating user information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Logic for updating user information
	// Implement your logic to update user information with the data from the request body
}

// DeleteUser handles deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Logic for deleting a user by ID
	// Implement your logic to delete a user by ID from a database or any source
}

// GetUser retrieves user information by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Logic for retrieving user information by ID
	// Implement your logic to retrieve user information by ID from a database or any source
}
