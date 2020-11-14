package main
import "fmt"

func main() {

		i:= 12
		switch {
		case i<= 15 :
			fmt.Println("the number less than or equal to fifteen")
			fallthrough
		case i >= 20 :
			fmt.Println(" the number is greater than or equal to twenty")		
			fallthrough
		default :
		  fmt.Println( " not in our range")
		}

}

