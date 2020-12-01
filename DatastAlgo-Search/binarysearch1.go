package main

import (
	"fmt"
)


func main(){
	a := []int{3,4,6,4,7,8,9,11,27,96,36}

	fmt.Println(binarysearch(a,0,len(a)-1,96))
}
func binarysearch(b[]int,l,h,v int)bool{
		if l>h{
			return false
		}

		mid := l+(h-l)/2
		if v==b[mid]{
			return true 
		} else if v>b[mid]{
			return binarysearch(b, mid+1,h,v)
		} else{
			return binarysearch(b,l,mid-1,v)
		}
}