package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
	"time"
)

func CommentsHanlder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllComments(w, r)
	case "POST":
		createComment(w, r)
	case "DELETE":
		deleteComment(w, r)
	}
}

func getAllComments(w http.ResponseWriter, r *http.Request) {
	//parsing
	var commentsData []models.Comments
	byteData,_ := os.ReadFile("db/comments.json")
	json.Unmarshal(byteData, &commentsData)

	//giving response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commentsData)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comments
	json.NewDecoder(r.Body).Decode(&newComment)


	var commentsData []models.Comments
	byteData,_ := os.ReadFile("db/comments.json")
	json.Unmarshal(byteData, &commentsData)

	for i := 0; i < len(commentsData); i++ {
		if commentsData[i].Id == newComment.Id {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Comment with such kind of ID is already exist")
			return
		}
	}
	newComment.CreatedAt = time.Now()
	commentsData = append(commentsData, newComment) 
	
	res,_ := json.Marshal(commentsData)
	os.WriteFile("db/comments.json", res,0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Successfully Created")
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	var deleteComment models.Comments
	json.NewDecoder(r.Body).Decode(&deleteComment)

	var commentsData []models.Comments
	bytedata,_ := os.ReadFile("db/comments.json")
	json.Unmarshal(bytedata, &commentsData)

	for i := 0; i < len(commentsData); i++ {
		if commentsData[i].Id == deleteComment.Id {
			commentsData = append(commentsData[:i],commentsData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "no comment with such kind of ID")
			return
		}
	}

	res,_ := json.Marshal(commentsData)
	os.WriteFile("db/comments.json", res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Successfully Deleted")
}