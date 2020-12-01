+
package main
import "fmt"
const arraySize = 26
type node struct{

	children [arraySize] *node
	isEnd bool
}
type tries struct{
	root *node
}
f.unc initTrie() *tries{
	result := &tries{root: &node{}}
	return result
}
func(t *tries) insert(w string){
	wlen := len(w)
	currentNode := t.root
	for i:=0;i<wlen;i++{
		chIndex := w[i]-'a'
		if currentNode.children[chIndex]==nil{
			currentNode.children[chIndex]=&node{}
		}
		currentNode= currentNode.children[chIndex]
	}
	currentNode.isEnd = true
}
func(t tries) search(w string) bool {
	l := len(w)
	currentNode := t.root
	for i:=0;i<l;i++{
		charIndex := w[i]-'a'
		if currentNode.children[charIndex] == nil{
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
		if currentNode.isEnd == true{
			return true
		}
	return false	
}


func main(){
	nameTrie := initTrie()
	
	nameList := []string{
			"raja",
			"raj",
			"david",
			"dravid",
			"arun",
			"aravind",
			"balu",
			"bala",
	}
	for _,v:= range nameList {
		nameTrie.insert(v)
	}
	fmt.Println(nameTrie.search("balu"))
	fmt.Println(nameTrie.search("aravindan"))
}