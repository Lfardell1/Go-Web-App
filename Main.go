package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lfardell1/Go-Web-App-Blog/Helpers"
	Route "github.com/lfardell1/Go-Web-App-Blog/Routes"
)

func main() {

	// start the router

	r := mux.NewRouter()

	// Establish the logger
	infoLog, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer infoLog.Close()

	errorLog, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer errorLog.Close()
	// Allow access to the static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	logger := Helpers.NewLogger()

	// View Endpoints
	r.HandleFunc("/", http.HandlerFunc(Route.IndexHandler))
	r.HandleFunc("/login", http.HandlerFunc(Route.LoginHandler))
	r.HandleFunc("/register", http.HandlerFunc(Route.RegisterHandler))
	r.HandleFunc("/logout", http.HandlerFunc(Route.LogoutHandler))
	r.HandleFunc("/create", http.HandlerFunc(Route.CreateHandler))
	r.HandleFunc("/update", http.HandlerFunc(Route.UpdateHandler))
	r.HandleFunc("/delete", http.HandlerFunc(Route.DeleteHandler))
	r.HandleFunc("/user", http.HandlerFunc(Route.UserHandler))
	r.HandleFunc("/user/posts", http.HandlerFunc(Route.UserPostsHandler))
	r.HandleFunc("/user/posts/:id", http.HandlerFunc(Route.UserPostHandler))
	r.HandleFunc("/posts", http.HandlerFunc(Route.PostsHandler))
	r.HandleFunc("/posts/:id", http.HandlerFunc(Route.PostHandler))
	r.HandleFunc("/about", http.HandlerFunc(Route.AboutHandler))

	// API Endpoints
	r.HandleFunc("/api/users", http.HandlerFunc(Route.GetUsers))
	r.HandleFunc("/api/posts", http.HandlerFunc(Route.GetPosts))
	r.HandleFunc("/api/posts/:id", http.HandlerFunc(Route.GetPost))
	r.HandleFunc("/api/posts", http.HandlerFunc(Route.CreatePost))
	r.HandleFunc("/api/users", http.HandlerFunc(Route.CreateUser))
	r.HandleFunc("/api/login", http.HandlerFunc(Route.Login))
	r.HandleFunc("/api/logout", http.HandlerFunc(Route.Logout))
	r.HandleFunc("/api/register", http.HandlerFunc(Route.Register))
	r.HandleFunc("/api/delete", http.HandlerFunc(Route.DeletePost))
	r.HandleFunc("/api/update", http.HandlerFunc(Route.UpdatePost))
	r.HandleFunc("/api/updateUser", http.HandlerFunc(Route.UpdateUser))
	r.HandleFunc("/api/deleteUser", http.HandlerFunc(Route.DeleteUser))
	r.HandleFunc("/api/getUser", http.HandlerFunc(Route.GetUser))
	r.HandleFunc("/api/getUserPosts", http.HandlerFunc(Route.GetUserPosts))
	r.HandleFunc("/api/getUserPosts/:id", http.HandlerFunc(Route.GetUserPost))

	// Allow logging on server
	writeLog := io.MultiWriter(os.Stdout, infoLog)
	logger.LogInfo("Server started on port 8080")
	logger.LogInfo("Press CTRL+C to exit")
	logger.LogInfo("Visit http://localhost:8080 to view the application")

	// Start the server with gorilla handler for logging

	http.ListenAndServe(":8080", handlers.LoggingHandler(writeLog, r))

}
