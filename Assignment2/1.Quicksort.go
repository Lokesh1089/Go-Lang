package main
import "fmt"

func main(){
	  s := []int{4,9,4,6,3,4,5,2,7}
	 fmt.Println("Before quick sort ->", s)

	 k := qsort(s)
	 fmt.Println("After quick sort ->" ,k)
}
func qsort(b []int)([]int){
	len := len(b)
	if len<1{
		return b
	}
	piv := b[0]
	var ls []int
	var rs []int

	for _,v := range b[1:] {
		if v<=piv{
			ls = append(ls, v)
		}
	}
	for _,v := range b[1:]{
		if v>piv{
			rs= append(rs,v)
		}
		
	}
	ls = qsort(ls)
	rs = qsort(rs)

	ls = append(ls, piv)
	ls = append(ls, rs...)
	return ls

}