package main 
import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/cuckoo", handler2)

	fmt.Println("server is starting....")
	http.ListenAndServe(":8080", nil)
}

func handler( w http.ResponseWriter, r *http.Request ){

	fmt.Fprintf(w, "hello everyone\n")

	fmt.Fprintf(w, "lets play with cuckoo..\n")
	
}

func handler2(w http.ResponseWriter, r *http.Request ){

	fmt.Fprintf(w, "cuckoo cuckooo !!!!!\n")
}