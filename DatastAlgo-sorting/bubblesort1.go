package main
import "fmt"

func main(){

	arr  :=  []int{2,5,7,1,3}

	fmt.Println("before sorting", arr)
	
	bubsort(arr)

}
func bubsort(arr[]int){
	  la := len(arr)
	  
	  for i:=0; i<la-1; i++{
		  for j:=0; j<la-i-1;j++{
			  if arr[j]>arr[j+1]{
				  arr[j], arr[j+1]= arr[j+1], arr[j]
			  }
		  }
	  }

	  fmt.Println("after sorting",arr)
	  
}