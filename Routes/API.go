package Routes

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lfardell1/Go-Web-App-Blog/Database"
	Helpers "github.com/lfardell1/Go-Web-App-Blog/Helpers"
	"github.com/lfardell1/Go-Web-App-Blog/models"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RequestHandler struct{}

func init() {

}

func GetTime(w http.ResponseWriter, r *http.Request) {
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

func PaginationHandler(w http.ResponseWriter, r *http.Request) {

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

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

// Login LoginAPIHandler handles the login functionality
func Login(w http.ResponseWriter, r *http.Request) {
	// Logic for handling user login, authentication, and generating tokens
	// Implement your login logic here, handle user authentication, and generate tokens if successful
	// TODO: Santise the form values
	email := r.FormValue("login_email")

	password := r.FormValue("login_password")

	log.Println(r.FormValue(email + " " + password))

	// create a user object

	var loggedUser models.User

	loggedUser.Password = password
	loggedUser.Email = email

	loggedUser, err := Database.LoginUser(loggedUser)

	log.Println(loggedUser)

	if err != nil {
		renderErrorMessage(w, "Error logging in")
		return
	}

	renderSuccessMessage(w, "User Created")

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
	// TODO: Santise the form values
	//check form values being submitted
	log.Println(r.FormValue("username") + " " + r.FormValue("email") + " " + r.FormValue("name") + " " + r.FormValue("password"))
	// Get the form values
	username := r.FormValue("username")
	email := r.FormValue("email")
	name := r.FormValue("name")
	password := r.FormValue("password")

	// Validate form values
	if username == "" || email == "" || password == "" || name == "" {
		renderErrorMessage(w, "Error creating user, Some fields weren't filled out!")
		return
	}

	// Create a new user model
	var NewUser models.User
	NewUser.Username = username
	NewUser.Email = email
	NewUser.Name = name
	NewUser.Password = password

	// Now we submit to the database and allow it to make some checks and we'll get back a user
	_, err := Database.CreateUser(NewUser)
	if err != nil {
		renderErrorMessage(w, "Error creating user")
		return

	}

	renderSuccessMessage(w, "User Created")

	// Save the session

	// Redirect to the home page

}

func ValidateUsername(w http.ResponseWriter, r *http.Request) {
	UsernametoCheck := r.FormValue("username")
	fmt.Println(UsernametoCheck)

	if UsernametoCheck == "" || UsernametoCheck == " " {

		renderErrorMessage(w, "Please Enter A Username")
	}
	if Database.CheckIfUsernameExists(UsernametoCheck) {
		// Render error message in the HTML itself
		errorMessage := "Username already exists"
		renderErrorMessage(w, errorMessage)
	}
}

func ValidateName(w http.ResponseWriter, r *http.Request) {
	NameToCheck := r.FormValue("name")
	fmt.Println(NameToCheck)
	if NameToCheck == "" || NameToCheck == " " {
		renderErrorMessage(w, "Please Enter A Name")
	}
	// might need to check names later who knows
}
func ValidateEmail(w http.ResponseWriter, r *http.Request) {

	EmailToCheck := r.FormValue("email")
	fmt.Println(EmailToCheck)
	if !IsValidEmail(EmailToCheck) {
		// Render error message in the HTML itself
		errorMessage := "Only alphabetic characters and no attempting escape characters"
		renderErrorMessage(w, errorMessage)
	}
	if Database.CheckIfEmailExists(EmailToCheck) {
		// Render error message in the HTML itself
		errorMessage := "You're already in our database!"
		renderErrorMessage(w, errorMessage)
	}
	fmt.Println(EmailToCheck)
}
func ValidatePassword(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm-password")
	if len(confirmPassword) < 7 {
		// Render error message in the HTML itself
		errorMessage := "Password less then 8 characters"
		renderErrorMessage(w, errorMessage)
	}
	if len(password) > 50 {
		// Render error message in the HTML itself
		errorMessage := "Password should be less then 50 characters"
		renderErrorMessage(w, errorMessage)
	}

}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func renderErrorMessage(w http.ResponseWriter, message string) {
	// Render the error message directly within the HTML
	errorHTML := `<p class="error-message  fade-me-in fade-me-out" style="transition: all 5s ease-out;" _"on load transition my *opacity to 0">` + message + `</p>`
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(errorHTML))
}
func renderSuccessMessage(w http.ResponseWriter, message string) {
	// Render the error message directly within the HTML
	errorHTML := `<section class="success-message" _="on load transition my *opacity to 0"> <p style="transition: all 800ms ease-in;" _="on click transition my *display to 'none'">` + message + `</p>`
	errorHTML += `<span style="font-size:10px;"></span></section>`
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(errorHTML))
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
