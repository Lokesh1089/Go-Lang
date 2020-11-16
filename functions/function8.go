package main 
import "fmt"

func main () {

  d := divide(5.0 , 8)

  fmt.Println("division of number is", d)
}

func divide(a,b float64) (res float64) {

	res = a/b
	if b == 0.0 {
		panic("cannot provide second value as zero")
	}
	res = a/b
	return
}