package main
import "fmt"

type stackPro struct{
	numList [] int
}

func(s *stackPro) push( n int){
	s.numList= append(s.numList,n)
}
func(s *stackPro) pull(){
	l := len(s.numList)-1
	s.numList = s.numList[:l]
}

func main(){
	stackList := stackPro{}

	nums := []int{66,8,75,32,41,16,9,7,3,96,100}

	for _,v := range nums {
		stackList.push(v)
	}
	
	fmt.Println(stackList)
	for i:=0;i<6;i++{
	stackList.pull()
	}
	fmt.Println(stackList)

}