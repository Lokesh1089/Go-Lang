package main
import "fmt"

type nodeN struct{
	name string
	next *nodeN
}
type hList struct {
	array[] *nodeN
	size int
}

func(h *hList) ignite(){
	h.size =100
	h.array = make([]*nodeN,h.size)

}
func(h *hList) magicfunction(name string) int{
	a := name
	sum :=0 
	for _,v := range a {
		sum = sum+int(v)
	}
	return sum%h.size
}
func(h *hList) ins(s string){
	index := h.magicfunction(s)
	newNode := new(nodeN)
	newNode.name = s
	newNode.next= h.array[index]
	 h.array[index] = newNode
}
func(h *hList) getAll(){
	for i:=0;i<h.size;i++{
	head := h.array[i]
	if head !=nil{
		fmt.Print("\nindex of array ",i)
	}
	if head!=nil{
		fmt.Print(" <- Name -> ", head.name, " " )
	}
	}
}

func main(){
	studentList := new(hList)
	studentList.ignite()
	studentList.ins("raghav")
	studentList.ins("harini")
	studentList.ins("dhakshaya")
	studentList.ins("rithik")
	studentList.ins("vidhya sri")
	studentList.ins("gokula priyan")
	studentList.ins("kanmani")
	studentList.ins("krithick varun")
	studentList.ins("hari")
	studentList.ins("siva sri")
	studentList.ins("mugilazhgan")

	studentList.getAll()

}