package main
import "fmt"

type nodePro struct{
	key  int
	next *nodePro
}
type hashTab struct {
	 hashArray[] *nodePro
	 size int
}
func(h *hashTab) iniz(){
	h.size =7
	h.hashArray = make([]*nodePro, h.size)
	for i:=0;i<h.size;i++{
		h.hashArray[i]=nil
	}
}

func(h *hashTab) hashMagic(k int) int{
	value := k
	return value % h.size
}
func(h *hashTab) insert(v int){
	idx := h.hashMagic(v)
	tempNewNode := new(nodePro)
	tempNewNode.key = v
	tempNewNode.next = h.hashArray[idx]
	h.hashArray[idx] = tempNewNode
}
func(h *hashTab) search(s int)bool{
	idx := h.hashMagic(s)
	head := h.hashArray[idx]
	for head !=nil{
		if head.key==s{
		return true
		}
		head = head.next
	}
	return false
}
func(h *hashTab) delete( d int) bool{
	idx := h.hashMagic(d)
	var neuNode *nodePro
	var head *nodePro 
	head = h.hashArray[idx]
	if head !=nil && head.key ==d{
		h.hashArray[idx] = head.next
		return true
	}
	for head!=nil{
		neuNode = head.next
		if  neuNode!=nil&&neuNode.key==d{
			head.next = neuNode.next
			return true
			}
		head = neuNode
	}
	return false
}
func(h *hashTab) print(){

	
	for i:=0;i<h.size;i++{
		head := h.hashArray[i]
		if head !=nil{
			fmt.Print("\n value at index ",i," is ")
		}
		for head!=nil{
			fmt.Print(head.key," ")
			head = head.next
		}
	}
}
func main(){
	ht := new(hashTab)
	ht.iniz()
	numList:=[]int{42,33,12,7,69,45,39,15,6}
	for _,v := range numList{
	ht.insert(v)
	}
	ht.print()
	fmt.Println("Search Result :", ht.search(12))
	fmt.Println("Delete Result :", ht.delete(69))
	ht.print()

}