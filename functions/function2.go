package main 
import "fmt"

func main() {


	for i:= 0 ; i < 5; i++ {
		sharemsg( " hello world", i)
	
	}
}

func sharemsg(msg string , idx int) {

 fmt.Println(msg)
 fmt.Println("the value of index is ",idx)
.
}