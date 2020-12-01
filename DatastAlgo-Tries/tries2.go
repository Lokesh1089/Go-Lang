package main
import "fmt"

const alphabet = 26
type nodePro struct{
	 child [alphabet]*nodePro
	 isEnd bool
}
type triePro struct{
	root *nodePro
}
func iniz()*triePro{
	return &triePro{root: &nodePro{}}
}

func(t *triePro) insert(word string){
	lenOfWord:= len(word)
	currNode := t.root
	for i :=0;i<lenOfWord;i++{
		charIx := word[i]-'a'
		 if currNode.child[charIx] == nil{
			currNode.child[charIx] = &nodePro{}
		 }
		 currNode = currNode.child[charIx]
	}
   currNode.isEnd = true
}

func(t *triePro) find(w string) bool{
	lw := len(w)
	currNode := t.root
	for i:=0;i<lw;i++{
		ci := w[i]-'a'
		if currNode.child[ci] == nil {
			return false
		}
		currNode = currNode.child[ci]
	}
	if currNode.isEnd  == true{
		return true
	}
return false
}

func main(){
	listTrie := iniz()
	list := []string{
		"egg",
		 "eat",
		 "mug",
		 "water",
		 "bath",
		 "mugs",
		 "sell",
		 "sold",
		 "bathsoap",
		 "eating",
		 "cut",
		 "cutcake",
		 "birthday",
	}

	for _,v := range list{
		listTrie.insert(v)
	}

fmt.Println("search result for word 'egg' is    ->", listTrie.find("egg"))
fmt.Println("search result for word  'water' is ->", listTrie.find("water"))
fmt.Println("search result for word 'happy' is  ->", listTrie.find("happy"))
fmt.Println("search result for word 'cake' is   ->", listTrie.find("cake"))
fmt.Println("search result for word 'bath' is   ->", listTrie.find("bath"))
fmt.Println("search result for word 'wet' is    ->", listTrie.find("wet"))

}