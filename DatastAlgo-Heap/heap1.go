package main
import "fmt"

type heap struct{
	array [] int
}

func(h *heap) insert(value int){
	h.array = append(h.array, value)
	h.heaping(len(h.array)-1)
}
func(h *heap) heaping(index int){
	for h.array[parent(index)]< h.array[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}

}

func parent(i int)int{
	return (i-1)/2
}
func right(i int)int{
	return 2*i+2
}
func left(i int)int{
	return i*2+1
}
func (h *heap)  swap(i1,i2 int){
	h.array[i1], h.array[i2] = h.array[i2],  h.array[i1]
}
func (h *heap) extract()int{
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
func(h *heap) maxHeapD(index int){
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
	m := heap{}
	bh := []int{10,20,30,5,7,9,11,13,15,17}
	for _,v := range bh {
		m.insert(v)
		fmt.Println(m)
	}
	for  i:=0;i<5;i++{
		m.extract()
		fmt.Println(m)
	}
}