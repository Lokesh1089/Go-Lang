package main
import "fmt"

type Stack struct{

	nums []int
}
// Pushing elements to slices
func(s *Stack) Push(i int){

	s.nums = append(s.nums, i)
}

//Last In Last Out
func(k *Stack) Pop(){
	l := len(k.nums)-1

	k.nums = k.nums[:l]


}


func main(){
	myStack := Stack{}
	fmt.Println("Before Push ->",myStack)
	myStack.Push(50)
	myStack.Push(150)
	myStack.Push(200)
	fmt.Println("After Push ->",myStack)
	myStack.Pop()
	fmt.Println("After Pop ->",myStack)


}