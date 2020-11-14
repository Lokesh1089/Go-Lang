// operators

package main 

import "fmt"

func main(){

  number := 197
   value := 45
	if value >1 && value < 100{
		fmt.Println("the value must be between 1 and 100")
	}
	if value > 100 || number > 100 {
		fmt.Println( "either value or number greater than 100")
	}

   if  value < number {
	   fmt.Println( "number greater than value")
   }

if  value > number {
	fmt.Println( "value is greater than number")	
	
}
if value == number {

	fmt.Println(" equal equal !")
}
 fmt.Println(number<=value, number>=value, number!=value)
 fmt.Println(false)
 fmt.Println(!false)
 fmt.Println(!true)
}