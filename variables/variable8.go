// variable datatype conversion string 
// "strconv"
 package main

 import ("fmt"
		"strconv"
 )
 func main(){

	var k int = 89
	fmt.Printf("%v, %T\n", k, k)

	var f string
	f = strconv.Itoa(k)
	fmt.Printf("%v, %T \n", f,f)

	var q int = 12
	fmt.Printf("%v, %T\n", q, q)
	
	var h string 
	h = strconv.Itoa(q)
	fmt.Printf("%v, %T\n", h, h)
 }