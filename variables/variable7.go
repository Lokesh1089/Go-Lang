// Converting datatype of variable 

package main

import "fmt"

func main (){

	var i float32 = 66.12
	fmt.Printf("%v, %T\n" , i, i)

	var j int 
	j = int (i)
	fmt.Printf("%v, %T \n", j , j)

	var x float32 = 4.7689
	fmt.Printf("%v, %T\n", x, x)
	
	var y int
	y = int (x)
	fmt.Printf("%v, %T", y, y)
}