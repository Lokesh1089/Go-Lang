package main

import "fmt"

func main() {
	arr := []int{-125,12,984,85,966,459,-984,-357,-963,427,127,-11,635,}
	res := linsear(arr, -984)

	if res != 1{
		fmt.Println("Founded!", res)
	} else {
		fmt.Println("Not Found")
	}

}
func linsear(arr []int, num int) int {
	for _, v := range arr {
		if v == num {
			return num
		}
	}
	return 1
}
