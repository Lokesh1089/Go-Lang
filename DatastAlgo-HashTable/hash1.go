package main
import "fmt"

type node struct{
	value int
	next *node
}
type hashTable struct{
	array[] *node
	tabSize int
}
func(h *hashTable) init(){
	h.tabSize =10
	h.array = make([]*node, h.tabSize)
	
}
func(h *hashTable) hash(key int) int{
	hval := key
	return hval%h.tabSize
}
func(h *hashTable) add(value int){
	index := h.hash(value)
	temp :=new(node)
	temp.value = value
	temp.next = h.array[index]
	h.array[index]=temp
}
func(h *hashTable) find(x int)bool{
	index :=  h.hash(x)
	head := h.array[index]
	for head!=nil{
		if head.value==x {
			return true
		}
		head = head.next
	}
	return false
}
func(h *hashTable) print(){
	for i:=0;i<h.tabSize;i++{
		head:=h.array[i]
		if head !=nil{
			fmt.Print("\n value at index ",i," is ")
		}
		for head!=nil{
			fmt.Print(head.value," ")
			head = head.next
		}
	}
	fmt.Println()
}
func(h *hashTable) remove(r int)bool{
	index := h.hash(r)
	var newNode *node
	var head *node
	head = h.array[index]
	if head != nil && head.value== r {
	   h.array[index] = head.next
	return true
	}
	for head != nil{
		newNode = head.next
		if newNode!=nil&&newNode.value==r{
			head.next = newNode.next
			return true
		}
		head = newNode
	}
	return false
}

func main(){
	hastab := new(hashTable)
	hastab.init()
	hastab.add(52)
	hastab.add(18)
	hastab.add(10)
	hastab.add(40)
	hastab.add(25)
	hastab.add(30)
	hastab.print()
	fmt.Println()
	fmt.Println(" Finding a value ", hastab.find(40))
	fmt.Println(" Finding a value", hastab.find(91))
	fmt.Println(" Removing value", hastab.remove(40))
	hastab.print()

}