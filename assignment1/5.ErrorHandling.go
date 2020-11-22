
// error handling to perform divisions on  positive numbers only

package main

import (
	"errors"
	"fmt")



func divof(a , b int)(int, error){

	div := 0
	if a <= 0 {
		
		return 0 , errors.New("first number is less than or equal to zero ")
	} 
	if b <=0 {
		return 0 , errors.New("second  number is less than or equal to zero")

	} 

		div = a/b

 return div , nil	
 
}

func main(){

	division , err:= divof(16,0)

	if err != nil {

		fmt.Println(err)
	}
	
	fmt.Println("divison of given two number is :", division)

}