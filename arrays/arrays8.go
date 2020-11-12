package main

import "fmt"

func main(){

	var imax [3][3] int 
	imax[0] = [3]int{1,0,0}
	imax[1] = [3]int{0,1,0}
	imax[2] = [3]int{0,0,1}

	fmt.Println( imax)
}