package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
type Student struct {
 ID string  `json:"id"`
 Name string `json:"name"`
 FatherName string `json:"fathername"`
 Class string `json:"class"`
}
var students[]Student = []Student{}

func main() {

r := mux.NewRouter()
students = append(students, Student{ID:"0", Name: "Kishore", FatherName: "Thennavan", Class: "10th-B"})
students = append(students, Student{ID:"1", Name: "Naveen", FatherName: "Raj", Class: "10th-C"})
r.HandleFunc("/students",getAll).Methods("GET")
r.HandleFunc("/students/{id}", getStudent ).Methods("GET")
r.HandleFunc("/students/{id}", postStudent ).Methods("POST")
r.HandleFunc("/students/{id}", updatePost ).Methods("PUT")
r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
log.Fatal(http.ListenAndServe(":8080", r))

}

func getAll(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/json")

json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/json")

parameters := mux.Vars(r)

for _,v := range students{
	if v.ID == parameters["id"]{
		json.NewEncoder(w).Encode(v)
		return
	}
}
json.NewEncoder(w).Encode(&Student{})
	
}

func postStudent(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")

var stud Student 
params := mux.Vars(r)
_ = json.NewDecoder(r.Body).Decode(&stud)
stud.ID = params["id"]
students = append(students, stud)
json.NewEncoder(w).Encode(students)
}
func deleteStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for kk ,v := range students {
		if v.ID == p["id"]{
			students = append(students[:kk], students[kk+1:]... )
			break
		} 
	}

	json.NewEncoder(w).Encode(students)

}


func updatePost(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
var stu string= mux.Vars(r)["id"]
id,err := strconv.Atoi(stu)
if err != nil{
	panic(err)
}
var updatestudent Student
_ = json.NewDecoder(r.Body).Decode(&updatestudent)
students[id] = updatestudent
json.NewEncoder(w).Encode(students)

}