// length of maps

package main 

import "fmt"

func main(){

	students := map[int] string {

		1: "vaibhav", 
		2: "nikhil",
		3: "anuraj",
		4: "vanishree",
	
	}
	fmt.Println(students)

	fmt.Printf("length : %v", len(students))
}

