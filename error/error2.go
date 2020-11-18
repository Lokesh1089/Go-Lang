
// error handling in square root

package main

import (
	"errors"
	"fmt"
	"math"
)

func main () {

	num , err := sq(625)

	if err != nil {
		fmt.Println("the error is ", err)
	}else {

		fmt.Println(num)
	}


}


func sq(a  float64)( float64,error) {

	num:= 0.0
	  if a <0 {
		  return 0 , errors.New("given number is negative")
	  } else {
  
		num = math.Sqrt(a)
		 
	  }

	  return num , nil
}