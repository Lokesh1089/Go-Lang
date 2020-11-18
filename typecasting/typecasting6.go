package main
import ("fmt"
		"strconv"
)

func main() {

	var a string = "68"
	x, _ := strconv.Atoi(a)

	fmt.Println(x)

	var b int = 68
	y := strconv.Itoa(b)

	fmt.Println(y)
}