package main
import "fmt"

func main() {

	students := []string { "athulya", "saravanan", "mithun", "maya", "velan"}

	fmt.Println(students)

	// 

	exceptmaya := append(students[:3], students[4:]...)
	
	fmt.Println(exceptmaya)

	fmt.Println(students)

}