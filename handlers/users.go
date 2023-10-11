package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
	"time"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("users ishladi")
	switch r.Method {
	case "GET":
		getAllUsers(w, r)
	case "POST":
		createUser(w, r)
	case "PUT":
		updateUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// req.Body parse qilamiz  | User
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	// File ni ochamiz, parse qilamiz | []User
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userdata)

	// yangi userni arrayga qo'shamiz

	for i := 0; i < len(userdata); i++ {
		if userdata[i].Username == newUser.Username {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "This username is already taken by someone")
			return
		}
	}

	newUser.Id = len(userdata)+1
	newUser.CreatedAt = time.Now().Format(time.RFC1123)
	newUser.UpdatedAt = time.Now().Format(time.RFC1123)
	userdata = append(userdata, newUser)

	// array ni faylga yozamiz
	res, _ := json.Marshal(userdata)
	os.WriteFile("db/users.json",res,0)

	// yangi userni jsonga o'zgaritirb responsega yozamiz | 
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userdata)

	json.NewEncoder(w).Encode(userdata)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// req.Body parse qilamiz  | User
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	// File ni ochamiz, parse qilamiz | []User
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
 	json.Unmarshal(byteData, &userdata)

	// yangi userni arrayga qo'shamiz
	for i := 0; i < len(userdata); i++ {
		if userdata[i].Id == newUser.Id {
			userdata[i].Username = newUser.Username
			userdata[i].Email = newUser.Email
			userdata[i].Age = newUser.Age
			userdata[i].UpdatedAt = time.Now().Format(time.RFC1123)
		}
	}

	// array ni faylga yozamiz
	res, _ := json.Marshal(userdata)
	os.WriteFile("db/users.json",res,0)

	// yangi userni jsonga o'zgaritirb responsega yozamiz | 
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(newUser)
	fmt.Fprint(w,"Updated Succesfully")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//parsing
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	// file to array variable
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userdata)

	for i := 0; i < len(userdata); i++ {
		if userdata[i].Id == newUser.Id {
			userdata = append(userdata[:i], userdata[i+1:]...)
		}
	}

	res, _ := json.Marshal(userdata)
	os.WriteFile("db/users.json",res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Deleted Succesfully")
}
