package main
import "fmt"

type heapPro struct{
	array[]int
}

func(h *heapPro) insert(v int){
	h.array = append(h.array, v)
	h.heapify(len(h.array)-1)
}

func(h *heapPro)heapify(indx int){
	for h.array[par(indx)]<h.array[indx]{
		h.swap(par(indx),indx)
		indx = par(indx)
		
	}
}
func par(i int) int {
	return (i-1)/2
}
func(h *heapPro) swap(k1 , k2 int){
	h.array[k1], h.array[k2]=h.array[k2], h.array[k1]
}

func right(i int) int {
	return i*2+2
}
func left(i int)int{
	return i*2+1
}
func (h *heapPro) extract()int{
	ex :=h.array[0]
	l := len(h.array)-1
	
if len(h.array)==0{
	fmt.Println("zero length array")
	return -1
}

h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapD(0)
	
	return ex
}
func(h *heapPro) maxHeapD(index int){
	lastIndex := len(h.array)-1
	l,r := left(index),right(index)
	childComp:=0
	for l<= lastIndex{
		if l==lastIndex{
			childComp=l
		}else if h.array[l]>h.array[r] {
			childComp = l
		} else{
			childComp=r
		}
		if h.array[index]<h.array[childComp]{
			h.swap(index,childComp )
			index = childComp
			l,r = left(index), right(index)
		} else{
			return
		}
	}
}

func main(){
	a := &heapPro{}
	b := []int{35,33,42,10,14,19,27,44,26,31}
	for _,v:=range b{
		a.insert(v)
		fmt.Println(a)
	}

	for i:=0;i<4;i++{
		a.extract()
		fmt.Println(a)
	}

}	