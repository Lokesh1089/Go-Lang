package main
import "fmt"

type list struct {
	 val string
	 right *list 
	 left *list 
}

func main() {

	ar := []string{"a","x", "d","e", "b", "c", "z"}
	b := &list{val:"f"}

	for _, v := range ar{
		b.ins(v)
	}

	b.ordered()
	fmt.Println()
	fmt.Println(b.find("e"))
	fmt.Println(b.find("k"))

}
func(l *list) ins(s string){
	if l.val<s{
			if l.right == nil{
				l.right = &list{val: s}
			} else{
				l.right.ins(s)
			}
	} else if l.val>s {
			if l.left == nil{
				l.left = &list{val: s}
			} else {
				l.left.ins(s)
			}
	}
}


func(l *list) ordered(){
	if l !=nil{
		l.left.ordered()
		fmt.Printf("%s  ", l.val)
		l.right.ordered()
	}
}

func(l *list) find(v string) bool{
	if l == nil{
		return false
	}
	if l.val>v{
		return l.left.find(v)
	} else if l.val<v{
		return l.right.find(v)
	}

	return true
}