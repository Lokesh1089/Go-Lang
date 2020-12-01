package main
import "fmt"

func main(){
	arr := []int{122,236,345,489,576,682,712,836,963}
	fmt.Println(binsear(arr,0,len(arr)-1,682))
}

func binsear(a[]int,lv,hv,sv int)bool{
	if lv>hv{
		return false
	}
	mid := lv+(hv-lv)/2
	if a[mid]==sv{
		return true
	} else if sv<a[mid]{
		return binsear(a,lv,mid-1,sv)
	}else {
		return binsear(a,mid+1,hv,sv)
	}
}