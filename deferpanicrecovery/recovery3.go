package main
import "fmt"

func main() {

	 fname := "nedumaaran"
	 //lname := "rajangam"
	 fullname(&fname , nil)

	

}

func fullname( fname *string, lname *string){


      defer func () {
		  if r := recover(); r != nil{
               fmt.Printf("function is recovered from panic: %s \n", r)
		  }
	  }()


	if fname == nil {
		panic("fname cannot be nil")
	} 
	if lname == nil {

		panic("last name cannot be nil")
} 

}