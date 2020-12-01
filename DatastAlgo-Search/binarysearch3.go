package main
import "fmt"

func main(){

	slice :=[]int{785,945,-42,44,-95,12,3,-450,100,150}
fmt.Println("unsorted..", slice)

	la := len(slice)
	  
	for i:=0; i<la-1; i++{					//sorting a unsoted slice by using bubble sort method
				for j:=0; j<la-i-1;j++{		//to perform binarysearch
			if slice[j]>slice[j+1]{
				slice[j], slice[j+1]= slice[j+1],slice[j]
			}
		}
	}
	var a[]int = slice

	fmt.Println("sorted ...",a)
	
	fmt.Println(bisearch(a,0,len(slice)-1,12))     //giving sorted slice as a input
}

func bisearch(arr[]int,low,high,sval int)bool{
	if low>high{
		return false
	} 
	mval := low+(high-1)/2
	if arr[mval]==sval{
		return true
	} else if sval>arr[mval]{
		return bisearch(arr,mval+1,high,sval)
	}else {
		return bisearch(arr,low,mval-1,sval)
	}
}