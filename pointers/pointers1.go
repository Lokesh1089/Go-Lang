package main 
import "fmt"

func main() {
 
	a := 35 
	b:=&a
	fmt.Println(a, *b)
	a = 12
	fmt.Println(a, *b)
	fmt.Println(&a, b)




}