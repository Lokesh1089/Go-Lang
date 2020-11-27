package main
import "fmt"

func main(){
	arr :=[]int{4,3,7,8,6,2,1}

	fmt.Println("initial array ->", arr)

	len := len(arr)

	for i:=1; i<len;i=i+2{   //checking even value for two condition
		if arr[i]<arr[i-1]{  // 1) is even num is less than previous element
			swap(arr, i, i-1)  // if yes then swap both element
		}
		if arr[i]<arr[i+1]{		//2) is even num is less than next element
			swap(arr, i , i+1)  // if yes then swap both element
		}
	}
	
		fmt.Println("zig zag matrix->",arr)

}

func swap(arr[]int, a,b int){
	 arr[a], arr[b] = arr[b], arr[a]
}