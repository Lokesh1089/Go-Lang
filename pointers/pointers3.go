package main 

import "fmt"
func main(){

	 a := [3] int{1,2,3}
	 b := a
	 fmt.Println( a, b)
	 b[1] = 44
	 fmt.Println(a,b)


	 c := []int{1,2,3}
	 d := c
	 fmt.Println(c,d)
	 d[0] = 77
	 fmt.Println(c,d)

	 stuMark := map[string]int {
					"bharathi": 477,
					"murali": 411,
					"naveen": 388,
				 }
	boys := stuMark
				 fmt.Println(stuMark, boys)

	boys["naveen"] = 453	
	fmt.Println(stuMark, boys)		 
}