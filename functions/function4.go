package main 
import "fmt"

func main() {

	sum(1,2,3,4,5,6,7)
}
func sum( values ... int) {
	 
	results := 0 
	for _,v := range values {

		results += v
	}

		fmt.Println(results)
}