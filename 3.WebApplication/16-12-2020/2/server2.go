package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
	Body  string
	Data  string
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("template.html")
	p := PageData{Title: "AK Books", Body: "Welcome to AK Book Store", Data: "List of books"}
	t.Execute(w, p)

}
