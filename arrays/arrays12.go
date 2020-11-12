package main
import "fmt"

func main(){

	 s := [6]int {4,5,6,7,8,9}
	fmt.Println(s)


	 t := &s
	 t[1] = 99
	 t[3] = 11
	 t[5] = 22
	 fmt.Println(t)
	 fmt.Println(t)

}