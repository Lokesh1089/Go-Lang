// head recursion

package main 
import "fmt" 

func num1(n int) {

	if n>0 {

		fmt.Println(n)

		num1(n-1)
	}
}

func main(){

	num1(5)
}