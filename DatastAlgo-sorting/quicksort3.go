package main
import "fmt"
func main(){

	sl:= []int{501,897,1,505, -356,-1,656,784,214,0,347,124,-908}

	fmt.Println("before :", sl)
	fmt.Println("After :", qsor(sl))
}

func qsor(s[]int) ([]int){
	len:= len(s)

	if len<2{
		return s
	}

	p := s[0]
	var l[] int
	var r[] int

	for _,v := range s[1:]{
	if v<=p{	
		 l= append(l, v)
		}
	}
	for _,v := range s[1:]{
	if v>p {	
		r= append(r, v)
		}
	}
	l = qsor(l)
	r = qsor(r)

	l = append(l, p)
	l= append(l, r...)

	return l

}