package main 

import "fmt" 

func main() {

	fmt.Println("ant")
panic(" no more calling from other function")
	animal()
}

func animal() {

	fmt.Println("bat")
	fmt.Println("cat")
panic("no more animals in the zoo ...")
	fmt.Println("dog")
	fmt.Println("elephant")
