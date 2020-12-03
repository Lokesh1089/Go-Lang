package main
import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

//Book struct 
type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}
//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
 var books[] Book

func main() {
	r := mux.NewRouter()

books= append(books, Book{ID:"1", Title:"Oru Puliya Marathin Kadhai", Author:&Author{Firstname:"Sundara", Lastname:"Ramasamy"}})
books= append(books, Book{ID:"2", Title:"Gopalla Gramam", Author:&Author{Firstname:"ki", Lastname:"Rajanarayanan"}})
books= append(books, Book{ID:"3", Title:"Ponniyin Selvan", Author:&Author{Firstname:"Kalki", Lastname:"Krishnamoorthy"}})

	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/books/{id}",getBook ).Methods("GET")
	r.HandleFunc("/books/{id}", postBook ).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getbooks(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	for _,v := range books{
		if v.ID == p["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
 json.NewEncoder(w).Encode(&Book{})
}
func postBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	p := mux.Vars(r)
	var boo Book
	_= json.NewDecoder(r.Body).Decode(&boo)
	boo.ID= p["id"]
	books = append(books,boo)
	json.NewEncoder(w).Encode(books)

}
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	p := mux.Vars(r)
	for k,v := range books{
		if v.ID== p["id"]{
			books = append(books[:k],books[k+1:]...)
		var boo Book
			_= json.NewDecoder(r.Body).Decode(&boo)
			boo.ID= p["id"]
			books = append(books,boo)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	p := mux.Vars(r)
	for k,v := range books{
		if v.ID== p["id"]{
			books = append(books[:k],books[k+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}