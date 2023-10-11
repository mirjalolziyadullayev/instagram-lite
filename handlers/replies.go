package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
	"time"
)

func RepliesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllReplies(w, r)
	case "POST":
		createReply(w,r)
	case "PUT":
		updateReply(w, r)
	case "DELETE":
		deleteReply(w, r)
	}
}

func getAllReplies(w http.ResponseWriter, r *http.Request) {
	//opening variable for copying data from file (Parsing)
	var repliesData []models.Replies
	// reading json file and saving to variable
	byteData, _ := os.ReadFile("db/replies.json")
	// converting json to variable []models.Replies
	json.Unmarshal(byteData, &repliesData)

	//sending taken result to client
	json.NewEncoder(w).Encode(repliesData)
}

func createReply(w http.ResponseWriter, r *http.Request) {
	//parsing r.Body
	var newReply models.Replies
	json.NewDecoder(r.Body).Decode(&newReply)

	//parsing json
	var repliesData [] models.Replies
	byteData, _ := os.ReadFile("db/replies.json")
	json.Unmarshal(byteData, &repliesData)

	for i := 0; i < len(repliesData); i++ {
		if repliesData[i].Id == newReply.Id {	
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "comment with such kind of ID is already exist")	
			return
		}
	}
	newReply.CreatedAt = time.Now()
	newReply.UpdatedAt = time.Now()
	repliesData = append(repliesData, newReply)
	//array variable to json db file
	res, _ := json.Marshal(repliesData)
	os.WriteFile("db/replies.json",res,0)

	//sendiing response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created Successfully")
}

func updateReply(w http.ResponseWriter, r *http.Request) {
	//parsing r.body from request to variable type models.Replies
	var newReply models.Replies
	json.NewDecoder(r.Body).Decode(&newReply)

	//parsing json 
	var repliesData []models.Replies
	byteData,_ := os.ReadFile("db/replies.json")
	json.Unmarshal(byteData, &repliesData)

	for i := 0; i < len(repliesData); i++ {
		if repliesData[i].Id == newReply.Id {
			repliesData[i].UpdatedAt = newReply.UpdatedAt
			repliesData[i].Text = newReply.Text
		}
	}
	//wrapping gotten data to db file json
	res, _ := json.Marshal(repliesData)
	os.WriteFile("db/replies.json", res,0)

	// sending response
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Updated Succesfully")
}

func deleteReply(w http.ResponseWriter, r *http.Request) {
	var deleteReply models.Replies
	json.NewDecoder(r.Body).Decode(&deleteReply)

	var repliesData []models.Replies
	byteData,_ := os.ReadFile("db/replies.json")
	json.Unmarshal(byteData, &repliesData)

	for i := 0; i < len(repliesData); i++ {
		if repliesData[i].Id == deleteReply.Id {
			repliesData = append(repliesData[:i], repliesData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "no reply with such kind of ID")
		}
	}

	res,_ := json.Marshal(repliesData)
	os.WriteFile("db/replies.json", res, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Deleted Successfully")
}