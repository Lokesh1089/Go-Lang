// odd or even

package main
import "fmt"

func main(){

 var num int 
 fmt.Printf("Enter a number :  ")
 fmt.Scanf( "%d", &num)
if num %2 == 0 {
	fmt.Printf("%v is even ", num)
} else {
	    fmt.Printf( "%v is odd", num)

}



}