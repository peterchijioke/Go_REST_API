package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// User is a struct that group all field into a single unit
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// post is a struct that group all field into a single unit
type post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []post = []post{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/add/", addItem).Methods("POST")
	http.ListenAndServe(":5000", r)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPost post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)
	json.NewEncoder(w).Encode(posts)
}
