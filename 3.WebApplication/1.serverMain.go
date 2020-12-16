package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandlerFunc)
	http.HandleFunc("/second", myHandlerFunc2)
	http.ListenAndServe(":8080", nil)

}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> Welcome To home page</h1>")
}
func myHandlerFunc2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "sample.html")
}
