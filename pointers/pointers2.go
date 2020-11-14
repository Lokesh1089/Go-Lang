package main 
import "fmt"

func main() {
  
	var a int 
	 a= 34
	 var b *int = &a
	 fmt.Println(a , *b)
		*b = 15
	 fmt.Println(a , *b)

}