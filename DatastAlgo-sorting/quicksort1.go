package main 
import (
	"fmt"
)
func main(){
	var slice = []int{12,1,-33,50,45,5,96,-4,22}
	fmt.Println(slice)
	slice = quicksort(slice)
	fmt.Println(slice)
}

func quicksort(slice[]int) ([]int){
	length := len(slice)
	if length<2{
		return slice
	}

	pivot :=slice[0]
	var left []int
	var right []int

	for _,v  := range slice[1:]{
		if v<=pivot{
			left= append(left, v)
		}
	}

	for _,v :=range slice[1:]{
		if v> pivot{
			right = append(right, v)
		}
	}

	right= quicksort(right)
	left= quicksort(left)

	left = append(left, pivot)
	left = append(left, right...)

	return left
}
