package main 
import(
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)
//MovieDetails struct holds a data about movie
type MovieDetails struct {
	MovieID string `json:"movieid"`
	MovieName string `json:"moviename"` 
	Actor string `json:"actor"`
	Actress string `json:"actress"`
	Director string `json:"director"`
	Producer *Producer
}
//Producer struct will help to group the production details
type Producer struct {
	ProductionComapny string `json:"productioncompany"`
	ProducerName string `json:"producername"`
}
var movies[]MovieDetails
func main(){
	router := mux.NewRouter()

	movies = append(movies, MovieDetails{MovieID:"35",MovieName:"Kushi",Actor:"Vijay", Actress:"Jothika",Director:"S.J.Surya",Producer:&Producer{ProductionComapny:"SriSuryaMovies", ProducerName:"A.M.Rathnam"} } )
	movies = append(movies, MovieDetails{MovieID:"36",MovieName:"KatradhuTamil",Actor:"jeeva", Actress:"Anjali",Director:"Ram",Producer:&Producer{ProductionComapny:"xxxx", ProducerName:"BaluMahendran"} } )

	router.HandleFunc("/movie",GetAllMovies).Methods("GET")
	router.HandleFunc("/movie/{movieid}", GetMovie).Methods("GET")
	router.HandleFunc("/movie/{movieid}",CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{movieid}", DeleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",router ))
}
//GetAllMovies function is used to retrive all movie details
func GetAllMovies(w http.ResponseWriter, r *http.Request){
json.NewEncoder(w).Encode(movies)
}
//GetMovie used to retrieve single movie data
func GetMovie(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	for _,v := range movies{
		if v.MovieID == params["movieid"]{
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	json.NewEncoder(w).Encode(&MovieDetails{})
}
//CreateMovie function used to create a new entry
func CreateMovie(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)
	var mov MovieDetails
	_ = json.NewDecoder(r.Body).Decode(&mov)
	mov.MovieID = params["movieid"]
	movies = append(movies,mov)
	json.NewEncoder(w).Encode(movies)


}
//DeleteMovie function is used to delete a entries
func DeleteMovie(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for k,v := range movies{
		if v.MovieID == params["movieid"]{
			movies = append(movies[:k], movies[k+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}