package main
import "fmt"

func main() {
  
	var num2 int
fmt.Printf("Enter a number: ")
fmt.Scanf("%v", &num2)

n := num2/2
if num2%2 == 0 {
	fmt.Printf("%v is even, %v times", num2, n)   // with divisible times
} else {
	fmt.Printf("%v is odd, %v times", num2, n)
}

}