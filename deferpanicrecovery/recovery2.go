package main 
import "fmt" 

func main() {

	 defer func() {
   
		if r:= recover(); r != nil {

			fmt.Printf("function is recovered from panic reason is: %s \n", r)
		}

	 }()

	 fmt.Println("hello")
	 fmt.Println("good morning")
	 panic("no more greeting for this day")

	 greet()
}

func greet() {

	fmt.Println("good evening")
	fmt.Println("good night")
}