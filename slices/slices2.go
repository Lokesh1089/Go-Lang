package main
import "fmt"

func main(){

   x := [] int {9,8,7,4,5}
	fmt.Println(x)
	
	y := x
	y[1] = 11
	fmt.Println(y)
	fmt.Println(x)

}