// match string


package main 
import(
	"fmt"
	"regexp"
)

func main() {

	a , err := regexp.MatchString("foo.", "seafood") 

	fmt.Println("result : ", a, err)

	b , err := regexp.MatchString("rocket.", "pistol raja")
	fmt.Println("result:", b , err)


	c , err := regexp.MatchString(".at",  "cat bat hat vat rat pat pot cot mat")

	fmt.Println("result:", c , err)

	d , err := regexp.MatchString(".s.",  "cat bat hat vat rat pat pot cot mat")

	fmt.Println("result:", d , err)

	exp, err := regexp.MatchString("hope", "hello everyone i hope you all doing good")
	fmt.Println("matching result:", exp , err)


}