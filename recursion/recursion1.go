// factorial


package main
import "fmt"

func main() {

   fmt.Println(factorial(5))

}

func factorial( n int) int {
 
	if n==1 || n ==0 {
		return n 
	} else {
		return n * factorial(n-1)
	}
   

}