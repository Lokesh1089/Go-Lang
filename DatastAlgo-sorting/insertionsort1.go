package main 
import "fmt"

func main(){

	arr := []int{8,4,1,3,7,5,2,6}

	insort(arr)
}

func insort(arr[]int){
	len := len(arr)

	for i:=1;i<len;i++{
		for j:=0;j<i;j++{
			if arr[j]>arr[i]{
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
fmt.Println("After sorting ->", arr)

}
