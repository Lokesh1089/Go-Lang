package main 
import "fmt"

func main() {

	s:= "hello bro!"
	for k,v := range s { //fmt.Println(k,v)
	fmt.Println( k , string(v))
	}
}