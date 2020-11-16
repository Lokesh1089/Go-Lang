package main
import "fmt"
func main() {
 
	b := mult(2,3,4,5)
	 fmt.Println(" multiplication of num is", b)
}

func mult(values...int)( res int ){

	res = 1
  for _,v := range values {
	res *= v
  	}
  return 
}