package main 
import "fmt"

func main (){

	var n complex64 = 2 + 5i
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T \n", imag(n), imag(n))
}