
//,ok 

package main 
import "fmt"

func main() {

	teacher := map[string]int{

		"vanmathi" : 12,
		"bhuvanesh": 13,
		"suresh": 14,
		"vandhana": 15,
		"priya" : 16,
		
	}

	fmt.Println(teacher)
	fmt.Println("physics",teacher["suresh"])
	fmt.Println( teacher["buvanash"]) // "bhuvanesh" 

	teach,ok := teacher["bhuvan"]  
	fmt.Println(teach, ok)


	tea,ok := teacher["vandhana"]
		fmt.Println(tea , ok)



}	