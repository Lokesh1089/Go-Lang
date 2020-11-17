package main 
import (

"fmt" 
"time"
) 

func main() {

 mychan1 := make(chan string)
 mychan2 := make( chan string)

 go func () {
	time.Sleep(1* time.Second )  
	mychan1 <- "hello everyone!"
	

 }() 

 go func () {
	time.Sleep( 5* time.Second) 
	mychan2 <- "welcome"


}() 

 for i :=1; i < 3; i++ {
	
	select {
		
	case hint1 := <- mychan1:
		fmt.Println(hint1)

	case hint2 := <- mychan2:
	     fmt.Println(hint2)
	}
 
 }

}