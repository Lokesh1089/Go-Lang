package main
import (
	"fmt"
	"log"
	"net/url"
)
func main(){
	ur , err := url.Parse("https://github.com/Lokesh1089/Go-Lang/tree/master/Assignment2")
	if err != nil{
		log.Fatal()
	}
	fmt.Println("Scheme: ", ur.Scheme)
	fmt.Println("Host: ", ur.Host)
	fmt.Println("Path: ", ur.Path)

	}


}