package main 
import "fmt" 

func main () {

	 var b int
	 for b=0; b<10; b++{

		fmt.Println(fibonacci(b))

	 }
	

}

func fibonacci( a int ) int {

	if a ==0 {
		return 0
	}
	if a == 1 {
		return 1
	} else {
		return fibonacci(a-1) + fibonacci(a-2)
	}
}