package main 
import "fmt" 

func num1 ( n int) {
	if n>0 {

		num1(n-1)
		fmt.Println(n)
	}
}

func main() {
	num1(5)
}