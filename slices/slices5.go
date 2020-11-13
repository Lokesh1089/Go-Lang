package main

import "fmt"

func main(){

	a := make( []int, 3)

	fmt.Println(a)
	fmt.Printf("length : %v \n", len(a))
	fmt.Printf("capacity : %v \n", cap(a))

	b := make( []int, 5,9)

	fmt.Println(b)
	fmt.Printf("length : %v \n", len(b))
	fmt.Printf("capacity : %v \n", cap(b))
}