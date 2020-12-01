package main

import "fmt"

func main() {
	arr := []int{99, 78, 7, -758, 54, -68, 8, 758, -785, 32, 49}
	res := linsear(arr, 123)

	if res != 0 {
		fmt.Println("yeah! i found the number")
	} else {
		fmt.Println("sorry! number is not in this box")
	}

}
func linsear(arr []int, num int) int {
	for _, v := range arr {
		if v == num {
			return num
		}
	}

	return 0
}
