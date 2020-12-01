package main
import "fmt"
type queueMax struct{
	name [] string
}

func(q *queueMax) enq(s string){
	q.name = append(q.name,s)

}
func(q *queueMax) deq() string{
	 l := len(q.name)
	 rem := q.name[0]
	 q.name = q.name[1:l]
	 return rem	
}

func main(){
   
	medicalQueue := &queueMax{}
	customers := [] string{
		"Balu",
		"Kavi",
		"Arun",
		"Rajesh",
		"Santhosh",
		"Naveen",
		"Peter",
	}
	for _,v := range customers{
		medicalQueue.enq(v)
		fmt.Println("Customer names are noted In First Come order",medicalQueue.name)
	}
	fmt.Println()
	for i:=0;i<4;i++{
	fmt.Printf(" %s went to the counter to buy a medicines\n", medicalQueue.deq() )
	}
	fmt.Println()
		fmt.Println("Peoples in the Queue", medicalQueue.name )
	

}