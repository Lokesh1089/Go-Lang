package main 
import "fmt" 

func main() {

	var num1 int = 25
	var num2 int = 4 
	var div float64

	div = float64(num1)/float64(num2)

	fmt.Printf("division of two num is %v %T", div , div)
}