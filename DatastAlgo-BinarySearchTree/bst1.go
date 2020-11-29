package main

import (
	"fmt"
)


var count int
type node struct{
 key int
  left *node
  right *node
}
func(n *node) insert(k int){
	if n.key<k{
		if n.right==nil{
			n.right =&node{key: k}
		} else{
			n.right.insert(k)
		}
	} else if n.key>k{
		if n.left ==nil{
			n.left= &node{key: k}
		} else{
			n.left.insert(k)
		}
	}
}
func(n *node) search(k int)bool{
	count++
	if n==nil{
	return false
	}
if n.key<k{
  return n.right.search(k)
} else if n.key>k{
    return n.left.search(k)
}
return true
}
func(n *node) findmin()int{
	if n.left == nil{
		return n.key
	}
	return n.left.findmin()
}

func(n *node) findmax()int{
	if n.right ==nil{
		return n.key
	}
	return n.right.findmax()
}

func (n *node) PrintInorder() {

	if n == nil {

		return
	}

	n.left.PrintInorder()
	fmt.Printf("%v ",n.key)
	n.right.PrintInorder()
}
func (n *node) Delete(value int) {
	n.remove(value)
}

func (n *node) remove(value int) *node {

	if n == nil {
		return nil
	}

	if value < n.key {
		n.left = n.left.remove(value)
		return n
	}
	if value > n.key {
		n.right = n.right.remove(value)
		return n
	}

	if n.left == nil && n.right == nil {
		n = nil
		return nil
	}

	if n.left == nil {
		n = n.right
		return n
	}
	if n.right == nil {
		n = n.left
		return n
	}

	smallestValOnRight := n.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	n.key = smallestValOnRight.key
	n.right = n.right.remove(n.key)
	return n
}


func main(){
	tree := &node{key:124}
	tree.insert(41)
	tree.insert(150)
	tree.insert(15)
	tree.insert(84)
	tree.insert(115)
	tree.insert(450)
	tree.insert(3)
	tree.insert(19)
	tree.insert(79)
	tree.insert(193)
	tree.PrintInorder()
	fmt.Println()
	fmt.Println(tree.search(2))
	fmt.Println(tree.search(115))
	fmt.Println(count)

	fmt.Printf("min is %d \n", tree.findmin())
	fmt.Printf("max is %d \n", tree.findmax())

	tree.Delete(84)
	tree.PrintInorder()


}
