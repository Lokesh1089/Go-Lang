package main 

import "fmt"

type Student struct{

	name string
	major string 
	age int
}

type teacher struct {
	Student
	name string 
	subject string 
}

func main () {

	 stuteach := teacher{

		Student :  Student{ name: "balaji", major: "mechanical", age: 21,},
		name: "kumaran",
		subject: "manufacturing technology",
	 }


	 fmt.Println( stuteach)
}