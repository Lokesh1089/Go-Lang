package main 
import "fmt"

func main(){
	//Array
	var a [7] int 
	a[2] = 55
	a[6] = 66
	fmt.Println(a)
	fmt.Println(a[2:4])

	//Slice

	var b = make([]int, 5)
	b[1] = 99
	b[2]= 77

	c := b[:3]
	b = append(b, 55)
	b = append(b, 111)
	fmt.Println(b)
	fmt.Println(c)
}