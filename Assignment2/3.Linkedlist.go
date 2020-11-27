package main
import "fmt"

type node struct{
	data int
	next *node
}
type list struct {
	head *node
	length int
}

func ( l *list) ins(n *node){
	second :=l.head
	l.head =n
	l.head.next = second
	l.length++
	}

func (l list) printData(){
	toprint := l.head
	
	for l.length !=0{
		fmt.Printf("%d ",toprint.data)
		toprint = toprint.next
		l.length--
	}
	fmt.Printf("\n",)
}	
func (l *list) deletevalue(value int){

	if l.length ==0{
		return
	}
	if l.head.data==value{
		l.head = l.head.next
		l.length--
		return
	}
	preToDelete := l.head
	for preToDelete.next.data != value{
		if preToDelete.next.next == nil{
			return
		}
		preToDelete = preToDelete.next
	}
	preToDelete.next = preToDelete.next.next
	l.length--
}

func main(){
  mylist := list{}
  node1 := &node{data: 48}
  node2 := &node{data: 21}
  node3 := &node{data: 75}
  node4 := &node{data: 11}
  node5 := &node{data: 36}
  node6 := &node{data: 7}
  mylist.ins(node1)
  mylist.ins(node2)
  mylist.ins(node3)
  mylist.ins(node4)
  mylist.ins(node5)
  mylist.ins(node6)
  mylist.printData()
  mylist.deletevalue(36)
  mylist.printData()
  mylist.deletevalue(7)
  mylist.deletevalue(89)
  mylist.printData()
}