package main 
import (
	"fmt" 
	"errors"
	
)

func main(){

 sum , err := sumofnum(18, 11)
if err != nil{

	fmt.Println( " the error is", err)
}else {

	fmt.Println(sum)
}
	
}

func sumofnum( start , end int) (int, error) {

	sum := 0

	if start>end {

		return 0 , errors.New("starting num is greater than end ")
	}
	 for i := start ; i <= end ; i++{

		 sum += i 
	     }
	 return sum , nil
    
}