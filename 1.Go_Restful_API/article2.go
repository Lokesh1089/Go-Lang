package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Post struct
type Post struct{
Title string `json:"title"`
Body string `json:"body"`
Author User  `json:"author"`
}
//User struct
type User struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email string `json:"email"`
}

var posts []Post = []Post{}
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/posts", addItem ).Methods("POST")
	r.HandleFunc("/posts", getallPost ).Methods("GET")	
	r.HandleFunc("/posts/{id}",getPost ).Methods("GET")
	r.HandleFunc("/posts/{id}",updatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", patchPost ).Methods("PATCH")	
	http.ListenAndServe(":5050", r)
}

func addItem(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
var newPost Post
json.NewDecoder(r.Body).Decode(&newPost)
posts= append(posts, newPost)
json.NewEncoder(w).Encode(posts)
}
func getallPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func getPost(w http.ResponseWriter, r *http.Request){
	var idparams string = mux.Vars(r)["id"]
	id , err :=  strconv.Atoi(idparams)
	if  id >= len(posts){
		w.WriteHeader(404)
		w.Write([]byte("No Post found in specified ID"))
	}
   
if err !=nil{
	w.WriteHeader(400)
	w.Write([]byte("ID could not be converted to integer"))
//	return
	}

	post := posts[id]
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(post)

}
func updatePost(w http.ResponseWriter, r *http.Request){
	var idparams string = mux.Vars(r)["id"]
	id , err := strconv.Atoi(idparams)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to integer"))
	}
	if id >= len(posts){
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}
	 var updatePost Post
	 json.NewDecoder(r.Body).Decode(&updatePost)
	posts[id] = updatePost 
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(posts)
}

func patchPost(w http.ResponseWriter, r *http.Request){
	var idparams string = mux.Vars(r)["id"]
	id , err := strconv.Atoi(idparams)
	if err != nil{
		w.WriteHeader(400)
		w.Write([]byte("ID could not converted to integer"))
	}
	if id >= len(posts){
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		//return
	}

	post := posts[id]
json.NewDecoder(r.Body).Decode(&post)

w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)


}