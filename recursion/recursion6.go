package main
import "fmt"

func fhello(){

	fmt.Println("hello bro")
	fhello()
}

func main() {

	fhello()
}