package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageInfo struct {
	Title string
	Body  string
	Data  string
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/data", formDataHandler)
	http.ListenAndServe(":9000", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("formfile.html")
	p := PageInfo{Title: "Vellore Super Market", Body: "We are Here to satisfy Your daily needs", Data: "LogIn here !"}
	t.Execute(w, p)

}
func formDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Please try again !")
	} else {
		mail := r.FormValue("mail")
		pass := r.FormValue("pword")
		fmt.Fprintf(w, "You are logged in with  Mail Id %s And your Password is %s", mail, pass)

	}
}
