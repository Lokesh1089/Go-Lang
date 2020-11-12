package main 

import "fmt"

func main(){

	var imax [3][3] int = [3][3]int {[3]int {1,0,0}, [3]int {0,1,0}, [3] int {0,0,1}}
	fmt.Println(imax)

	fmt.Printf("length of array : %v\n", len(imax))
	fmt.Printf("imax #2: %v\n", imax[2])
}