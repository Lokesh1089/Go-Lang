package main
import "fmt"

func main() {

	 fname := "nedumaaran"
	 //lname := "rajangam"
	 fullname(&fname , nil)

}

func fullname( fname *string, lname *string){

	if fname == nil {
		panic("fname cannot be nil")
	} 
	if lname == nil {

		panic("error : last name cannot be nil")
} 

fmt.Printf("%s %s \n ", *fname ,*lname)

}