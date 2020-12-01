package main
import (
	"log"
	"net/url"
	"fmt"
)
func main() {
	link1 := "https://www.golangprograms.com/url-parser-in-golang.html"
	link2 := "https://www.eduonix.com/new_dashboard/Learn-Database-Design-using-PostgreSQL"

	u , err := url.Parse(link1)
	if err !=nil{
		log.Fatal()
	}
	u2 , err := url.Parse(link2)
	if err !=nil{
		log.Fatal()
	}
	fmt.Println("Scheme :", u.Scheme)
	fmt.Println("Host: ", u.Host)
	fmt.Println("Path: ", u.Path)

	fmt.Println("RawPath: ", u2.RawPath )
	fmt.Println("Query: ", u2.Query )
	fmt.Println("RawQuery: ", u2.RawQuery )
	fmt.Println("Port: ", u2.Port )
	fmt.Println("RawFragment: ", u2.RawFragment )
	fmt.Println("Fragment: ", u2.Fragment )


	  

}