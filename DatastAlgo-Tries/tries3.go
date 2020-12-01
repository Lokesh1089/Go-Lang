package main
import "fmt"

const size = 26

type nodeOne struct{
	childs[size] *nodeOne
	isClimax bool

}
type trieOne struct{
	root *nodeOne
}
func inis() *trieOne{
	return &trieOne{root: &nodeOne{}}
}
func(t *trieOne) add(s string){
	ls := len(s)
	cNode := t.root
	for i:=0;i<ls;i++{
		cin := s[i]-'a'
	if	cNode.childs[cin]==nil{
			cNode.childs[cin]=&nodeOne{}
		}
		cNode.childs[cin] = cNode
	}
	cNode.isClimax = true
}
func(t *trieOne) search(s string)bool{
	ls := len(s)
	curNode := t.root
	for i:=0;i<ls;i++{
		ci := s[i]-'a'
	if curNode.childs[ci] == nil{
		return false 
	}
	curNode = curNode.childs[ci]	
}
	if curNode.isClimax == true{
	  return true
	}
	return false 
}
func main(){
	
	alist := inis()
	aToz := [] string{
		"abc",
		"abb",
		"cdx",
		"cbs",
		"fff",
		"fdf",
	}
	for _,v:= range aToz{
		alist.add(v)
	}

	fmt.Println(alist.search("xyz"))
	fmt.Println(alist.search("abc"))
	fmt.Println(alist.search("fff"))
	fmt.Println(alist.search("klt"))

}