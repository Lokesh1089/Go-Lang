package main
import(

	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)
//Person is used to define the details
type Person struct{
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}
//Address used to store person's address
type Address struct{
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
var people[]Person

//GetPerson details from server
func  GetPerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _,v := range people{
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
json.NewEncoder(w).Encode(&Person{})

}
//GetPeople is Used to collect all details
func GetPeople(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)

}
//CreatePerson is Used to create a person
func CreatePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people,person)
	json.NewEncoder(w).Encode(people)
}
//DeletePerson is Used to detelete particular person  details
func DeletePerson(w http.ResponseWriter, r *http.Request){

params := mux.Vars(r)
for k,v:= range people{
	if v.ID==params["id"]{
		people = append(people[:k], people[k+1:]...)
		break
	}
}
json.NewEncoder(w).Encode(people)
}

func main(){
	router := mux.NewRouter()
	people = append(people,Person{ID:"1", Firstname: "Rajendra", Lastname:"Chozhan", Address:&Address{City:"GangaiKondaChozhaPuram", State:"Tamilnadu"} } )
	people = append(people,Person{ID:"2", Firstname: "Rajaraja", Lastname:"Chozhan", Address:&Address{City:"Thanjai", State:"Tamilnadu"} } )
	people = append(people,Person{ID:"3", Firstname: "velpaari", Address:&Address{City:"ParambuNadu"} } )
	
	router.HandleFunc("/people",GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":2122",router))
}