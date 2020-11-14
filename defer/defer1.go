package main 
import "fmt"

func main() {

	a := 45
	b := 12
defer	fmt.Println( a)
defer	fmt.Println( b)
	result := a+b
	fmt.Println(result)
}