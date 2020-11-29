package main
import "fmt"

type list struct{
	head *node
	tail *node
}
type node struct {
	value int
	next *node
	pre *node
}

func(l *list) first() *node{
	return l.head
}
func(l *list)last()*node{
	return l.tail
}

func(l *list) push(value int){
	n := &node{value: value}
	if l.head == nil{
		l.head =n
	} else{
		l.tail.next = n
		n.pre = l.tail
	}
	l.tail = n
}
func(n *node) nex() *node{
	return n.next
}
func(n *node) prev() *node{
	return n.pre
}
func(l *list) print(){
 n := l.head
 for n != nil{
	 fmt.Println(n.value)
	 n = n.next
 }

}

func main(){
 l := &list{}
 l.push(10)
 l.push(9)
 l.push(8)
 l.push(7)
 l.print()
fmt.Println("--------------------")
 a := l.first()
 for {
	 fmt.Printf("%v ",a.value)   //next
	 a = a.nex()
	 if  a==nil{
		 break
	 }
 }
fmt.Println()
b := l.last()
for {
	fmt.Printf("%v ",b.value)    //previous
	b = b.prev()
	if b==nil{
		break
	}
}


}