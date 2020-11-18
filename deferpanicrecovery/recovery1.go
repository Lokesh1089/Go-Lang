package main

import "fmt" 

func main() {

	fmt.Println("anand")
	remain()
}

func remain() {

	defer func() {

		if r := recover(); r != nil {

			fmt.Printf("The function is recovered from panic the reason is : %s \n", r)
		}
	}()

	fmt.Println("babu") 
	fmt.Println("chezhiyan") 
	panic("remaining students are left from school")
	fmt.Println("deva")
	fmt.Println("elevarasan")
}