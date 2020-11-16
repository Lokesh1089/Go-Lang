package main 
import "fmt"

func main() {
  
	a, err := div(4.0,7.0)
	if err != nil{
		fmt.Println(err)
		return
	}


	fmt.Println("division of num is", a)

}

func div(c, d float64)(float64, error) {

   if d == 0.0 {
	   return 0.0 , fmt.Errorf( " socond number cant be zero")
   }

   return c/d , nil

}