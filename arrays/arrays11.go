package main 

import "fmt"

func main() {

	k := [...]int{1,2,3,4,5}
	
	fmt.Println(k)
	fmt.Printf("length of %v \n", len(k) )
	fmt.Printf("capacity of %v \n", cap(k) )

	m := k
	fmt.Println(m)
	m[1] = 88
	fmt.Println(m)
	fmt.Println(k)
	
	q := &k
	q[1] = 99
	fmt.Println(q)
	fmt.Println(k)


}