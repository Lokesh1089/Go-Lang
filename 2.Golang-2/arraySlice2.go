package main
import "fmt"

func main(){

//Slice
 var a [] int 
//Array
  inputArray := [6] int {25,50,75,100,125,150}	

	for _,v := range inputArray {
	a = append(a, v)
	fmt.Println(a)
	}
 
}
