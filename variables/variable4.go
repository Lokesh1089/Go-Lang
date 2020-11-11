// variable shadowing 

package main 

import "fmt"

var(

	physicsLabIncharge string = "muthu kumar"
)

func main(){
  
	 physicsLabIncharge = " rajesh"
 
	fmt.Println("new lab incharge", physicsLabIncharge)

}
