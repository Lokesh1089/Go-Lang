package main
import "fmt"

func main(){

	a := []int{98,52,63,-44,74,-33,56,-11,36}
	fmt.Println("Before insertion sort", a)

	insertsort(a)
}

func insertsort(b[]int){
	lenthOfArray := len(b)

	for k :=1;k<lenthOfArray;k++{
		for v:=0; v<k;v++{
			if b[v]>b[k]{
				b[v],b[k] =b[k],b[v]
			}
		}
	}
	fmt.Println("After insertion sort",b)
}