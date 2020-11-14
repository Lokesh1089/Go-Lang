package main 

import "fmt"
func main() {

  fmt.Println("start")
  defer fmt.Println( "this was deffred")
	panic( "something bad happened")
	fmt.Println("end")
}