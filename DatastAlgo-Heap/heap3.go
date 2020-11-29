package main
import "fmt"

type studentList struct{
	name [] string
}

func(s *studentList)insert(a string){
	 s.name = append(s.name, a)
	 s.hefy(len(s.name)-1)

}
func(s *studentList) hefy(i int){
	for s.name[paren(i)]< s.name[i] {
		s.swap(paren(i),i)
		i = paren(i)
	}
}

func (s *studentList) swap(t1, t2 int){
	s.name[t1], s.name[t2] = s.name[t2], s.name[t1]
}

func paren(i int) int{
	return (i-1)/2
}

func main(){
	ls := &studentList{}
	nl := []string{"a","h","g","b","s","v","f","k"}

	for _,v :=range nl{
		ls.insert(v)
		fmt.Println(ls)
	}
}