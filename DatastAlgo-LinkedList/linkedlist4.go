package main
import "fmt"

type studentlist struct{
    head *nod
	tail *nod
}
type nod struct{
	name string
	mark int
	next *nod
	prev *nod
}

func(s *studentlist) first() *nod{
	return s.head
}

func(s *studentlist) last() *nod{
	return s.tail
}
func(s *studentlist) insert(a string, m int){
	b := &nod{name: a,
				mark: m,}

	if s.head == nil {
		s.head = b
	} else {
		s.tail.next = b
		b.prev = s.tail
	}			
	s.tail = b
}
func (n *nod) nextt() *nod{
	return n.next
}
func( n *nod) previous() *nod{
	return n.prev
}

func main(){

   ls := studentlist{}
   ls.insert("santhosh", 487)
   ls.insert("varun", 395)
   ls.insert("saranya",296)
   ls.insert("naveen", 434)
   ls.insert("bhuvana",456)

firststu := ls.first()
fmt.Println("first student is ", firststu.name)

laststu := ls.last()
fmt.Println("Last student  is ", laststu.name )

	a := ls.first()
	for {
		fmt.Println(a.name,"-",a.mark)
		a = a.nextt()
		if a==nil{
			break
		}
	}
fmt.Println("---------previous---------")
	b := ls.last()
	for {
		fmt.Println(b.name,"-", b.mark)
		b = b.previous()
		if b==nil{
			break
		}
	}
	
}
