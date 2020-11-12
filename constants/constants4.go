package main

import "fmt"

func main() {
	const a int = 13
	var b int = 18
	fmt.Println(a + b)

	const x = 11
	const y int16 = 22

	fmt.Printf("%v, %T \n", x, x)
	fmt.Println(x + y)
}
