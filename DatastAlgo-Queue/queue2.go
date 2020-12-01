package main
import "fmt"

type queuePro struct{
	nums [] int
}

func(q *queuePro) enqueue(n int){
	q.nums = append(q.nums, n)
}
func(q *queuePro) dequeue() int{
	l := len(q.nums)
	r := q.nums[0]
	q.nums = q.nums[1:l]
	return r
}
func main(){
	numList := &queuePro{}
	numbers :=[]int{25,35,45,55,100,200,300,400,500}
	for _,v := range numbers{
		numList.enqueue(v)
		fmt.Println("Queue getting increases",numList.nums)
	}
	for i:=0;i<5;i++{
		
		fmt.Printf("\n Token Number <- %d -> WentOut from the Queue" ,numList.dequeue())
	}
fmt.Println()
	fmt.Println("Remaining Token Numbers in the Queue ", numList.nums)

}