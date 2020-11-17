// indirect recursion

package main
import "fmt" 

func fone (n int ){

	if n>= 0 {

		fmt.Println("in first function:", n)

		ftwo(n-1)
	}
}

func ftwo( n int){

	if n >= 0 {
		fmt.Println("in second function :", n)

		fone(n-1)
	}
}


func main () {
   
	fone(10)
	ftwo(-1)

}