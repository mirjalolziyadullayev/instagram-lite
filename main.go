package main

import (
	"fmt"
	"instagram/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.GetHomePage)

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/posts", handlers.PostsHandler)
	http.HandleFunc("/comments", handlers.CommentsHanlder)
	http.HandleFunc("/replies", handlers.RepliesHandler)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
