package main 

import "fmt"

func main() {

	a := []int{}

	fmt.Println(a)
	fmt.Printf("length : %v \n", len(a))
	fmt.Printf("capacity : %v \n", cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("length : %v \n", len(a))
	fmt.Printf("capacity : %v \n", cap(a))

	a = append(a, 2,3,4,5,6,7,8)
	fmt.Println(a)
	fmt.Printf("length : %v \n", len(a))
	fmt.Printf("capacity : %v \n", cap(a))

	a = append(a, 1,2,3,4,5,6,7)
	fmt.Println(a)
	fmt.Printf("length : %v \n", len(a))
	fmt.Printf("capacity : %v \n", cap(a))


}