package main
import "fmt"

type stackString struct{
	name [] string
}
func(s *stackString)  push(n string){
	s.name = append(s.name,n)
}
func(s *stackString) pull(){
	length := len(s.name)-1
	s.name  = s.name[:length]
}
func main(){
	nameList := &stackString{}

	names :=[]string{"murugan","vivek","prabhu","vishnu","ram","kumar","kaalidas","nedumaran","ashwin"}
	fmt.Println(nameList)
	for _,v := range names{
		nameList.push(v)
		fmt.Println(nameList,"->")
	}

	for i:=0;i<len(names);i++{
		nameList.pull()
		fmt.Println(nameList,"<-")
	}
}