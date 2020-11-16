package main 
import "fmt"

func main() {

	a := sum(1,2,3,4,5,6,7)

	fmt.Println("the sum is", *a)
}
func sum(values ... int) *int {
	 
	results := 1 
	for _,v := range values {

		results += v
	}
	return &results

}