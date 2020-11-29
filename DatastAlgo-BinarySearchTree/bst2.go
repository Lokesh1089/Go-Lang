package main
import "fmt"

type onenode struct{
	value int
	left *onenode
	right *onenode
}

func(a *onenode) insert(v int){
	if v<a.value{
		if a.left == nil{
			a.left = &onenode{value: v}
		} else {
			a.left.insert(v)
		}
	} else if v>a.value{
		if a.right == nil{
			a.right = &onenode{value: v}
		} else {
			a.right.insert(v)
		}
	}
}

func(a *onenode) viewInorder(){
	if a !=nil{
	a.left.viewInorder()
	fmt.Printf("%v ", a.value)
	a.right.viewInorder()
	}
}
func (a *onenode) max() int{
	if a.right==nil{
		return a.value
	}
	return a.right.max()
}
func (a *onenode) min() int{
	if a.right==nil{
		return a.value
	}
	return a.left.min()
}

func main(){
	tr := &onenode{value: 80}
	arr := []int{75, 12,97,150,33,19,3,9,28,105}
	for _,v := range arr {
		tr.insert(v)
	}
	tr.viewInorder()
	fmt.Println()
	fmt.Printf("max value is %d \n", tr.max())
	fmt.Printf("min value is %d \n", tr.min())
	
}