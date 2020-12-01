package main
import "fmt"

func main() {

	a := []int{8,2,6,1,7,9,10,25,45,96,3,44,1}
   
	var c bool = linearSearch(a,25 )
	if c ==true {
		fmt.Println("searching result for the value 25 is ->", c )
	} else{
		fmt.Println("eror 404 value not found !")
	}
	
}

func linearSearch(a[]int, val int)bool {
	for _,v:= range a {
		if v==val{
			return true
		}
	}
	return false
}