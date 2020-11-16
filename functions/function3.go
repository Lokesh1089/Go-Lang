package main
import "fmt"

func main() {

	greet := "hello !"
	name := "varun"
	saygreeting(&greet , &name)
	//fmt.Println(name)
}

func saygreeting(greet,name *string ) {

	fmt.Println(*greet , *name)
	*name = "rajan"
	fmt.Println(*name)
}