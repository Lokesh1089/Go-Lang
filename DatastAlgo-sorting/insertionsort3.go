package main 
import "fmt"

func main(){
	array :=[]int{998,789,455,-975,234,-789,124,0,77,-505,606}
	insort(array)
}

func insort(arr[]int){
 	len := len(arr)
	for i:=1;i<len;i++{
		for j:=0;j<i;j++{
			if arr[j]>arr[i]{
				arr[j],arr[i] = arr[i], arr[j]
			}
		}
	}
	for _,v := range arr{
		fmt.Println(v)
	}
}