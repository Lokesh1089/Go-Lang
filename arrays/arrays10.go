package main

import "fmt"

func main () {

	a := [...]int{9,8,7,6}
	fmt.Println(a)
	
	b := a 
	fmt.Println(b)

	b[1] = 1
	fmt.Println(b)
}