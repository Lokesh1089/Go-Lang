package main
import "fmt"

//creating structure named as queue
type Queue struct{
	values [] int
}
//adding element to the slice
func ( q *Queue) Enqueue( a int){
	q.values = append(q.values, a)
}

//first in first out ->removing first element
func(q *Queue) Dequeue()int{
	ln := len(q.values)
	r := q.values[0]
	q.values = q.values[1:ln]
	return r
}


func main() {
	myQueue := Queue{}
	fmt.Println("Before Enqueue ->",myQueue)
	myQueue.Enqueue(25)
	myQueue.Enqueue(50)
	myQueue.Enqueue(75)
	myQueue.Enqueue(100)
	fmt.Println("After Enqueue ->",myQueue)
	myQueue.Dequeue()
	fmt.Println("After Dequeue ->",myQueue)


}