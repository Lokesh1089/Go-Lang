package main

import (
	"fmt"
	"net/http"
	"html/template"
)
//PageData input for our web page
type PageData struct{
	Title string
	Body string
}
func main(){

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/olduser", oldUserHandler)
	http.HandleFunc("/newuser",newUerHandler )
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	t , _ := template.ParseFiles("page.html")
	p := PageData{Title: "Madras Mens Shop", Body: "Welcome To Madras Mens Official Website !"}
	t.Execute(w,p)

}
func oldUserHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost{
		fmt.Fprintf(w, "Please try again...")
	} else {
		mail := r.FormValue("mail")
		pass := r.FormValue("pword")
		fmt.Fprintf(w, "You are logged in with  Mail Id %s And your Password is %s", mail, pass)
	}

}
func newUerHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		fmt.Fprintf(w, "Please try again...")
	} else {
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		DateOfBirth := r.FormValue("DateOfBirth")
		PhoneNumber := r.FormValue("PhoneNo")
		Password := r.FormValue("Password")
		Email := r.FormValue("Email")

		fmt.Fprintf(w, "You are Created Account with following details\n")
		fmt.Fprintf(w, "Name  %s %s \n", fname , lname)
		fmt.Fprintf(w, "DOB %s \n", DateOfBirth)
		fmt.Fprintf(w, "PhoneNo %s \n", PhoneNumber)
		fmt.Fprintf(w, "Email %s \n", Email)
		fmt.Fprintf(w, "Password %s \n", Password)
		}	
}