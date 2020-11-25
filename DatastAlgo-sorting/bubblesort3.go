package main
import "fmt"
func main(){

	ar :=[]int{328,258,-789,0,557,123,1,-325,237,-985,896}
	fmt.Println("Before Sorting ", ar)
	bsort(ar)
}
func bsort(ar[]int){
	lenar := len(ar)
	for x:=0;x<lenar-1;x++{
		for y:=0;y<lenar-x-1;y++{
			if ar[y]>ar[y+1]{
				ar[y], ar[y+1]=ar[y+1],ar[y]
			}
		}
	}
	fmt.Println("After Sorting", ar)
}

