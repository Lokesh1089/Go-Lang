package main

import (
	"fmt"	
	"math"
)

func main(){

	var a , b int = 13,18
	var c float64 = math.Sqrt(float64(a*a+ b*b))
	var z uint = uint(c)
fmt.Println(a,b,z)

}