package main
import (
	"fmt"
)

func main(){
	a :=[]int{122,45,-95,189,-202,0,176,550,20,14,60}
	fmt.Println("Before Sorting")
	for _,v:= range a{
	    fmt.Println(v)
	}
	selectsort(a)
}

func selectsort(b []int){
	lenOfArray:=len(b)

	for i:=0;i<lenOfArray-1;i++{

		for j:=1+i;j<lenOfArray;j++{
			if b[i]>b[j]{
				b[i], b[j] = b[j],b[i]
			}
		}
	}
	fmt.Println("After Sorting")
	for _,v := range b {

		fmt.Println(v)
	}
}