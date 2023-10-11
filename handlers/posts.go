package handlers

import (
	"encoding/json"
	"time"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllPosts(w, r)
	case "POST":
		createPost(w, r)
	case "PUT":
		updatePost(w, r)
	case "DELETE":
		deletePost(w, r)
	}
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	var postsData []models.Posts
	byteData,_ := os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)

	json.NewEncoder(w).Encode(postsData)
}
func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Posts
	json.NewDecoder(r.Body).Decode(&newPost)

	var postsData []models.Posts
	byteData,_:= os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)

	for i := 0; i < len(postsData); i++ {
		if postsData[i].Id == newPost.Id {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Post with such kind of ID already exist")
			return
		} 
	}
	
	newPost.Id = len(postsData)+1
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()
	postsData = append(postsData, newPost)

	res,_ := json.Marshal(postsData) 
	os.WriteFile("db/posts.json",res,0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Succesfully Created")
}
func updatePost(w http.ResponseWriter, r *http.Request) {
	var updatePost models.Posts
	json.NewDecoder(r.Body).Decode(&updatePost)

	var postsData []models.Posts
	byteData, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)

	for i := 0; i < len(postsData); i++ {
		if postsData[i].Id == updatePost.Id {
			postsData[i].UserId = updatePost.UserId
			postsData[i].Title = updatePost.Title
			postsData[i].Content = updatePost.Content
			postsData[i].UpdatedAt = time.Now()
		}
	}

	res, _ := json.Marshal(postsData)
	os.WriteFile("db/posts.json",res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Updated Successfully")
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	var deletePost models.Posts
	json.NewDecoder(r.Body).Decode(&deletePost)

	var postsData []models.Posts
	byteData, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)

	for i := 0; i < len(postsData); i++ {
		if postsData[i].Id == deletePost.Id {
			postsData = append(postsData[:i],postsData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w,"no posts with such kind of ID")
			return
		}
	}

	res, _ := json.Marshal(postsData)
	os.WriteFile("db/posts.json",res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Deleted Successfully")
}