package main 

import "fmt"

func main (){

	var students [3] string 
		students[0] = "monish"
		students[1] = "kiruba"
		students[2] = "muthu"

		fmt.Println( "students:", students)

		students[0] = "kumar"
		students[2] = "manick"

		fmt.Println("updated :", students)
}