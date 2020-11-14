package main 
import "fmt"

func main(){
  var num interface{} = 12
  switch num.(type) {
  case int:
	 fmt.Println("num is an int")
	break
	 fmt.Printf("value of num %v", num)
  case float64 :
	fmt.Println("num is float64")
  case string :
	fmt.Println("num is string")
default :
   fmt.Println("num is another type")			 
	}

}
