package main
import "fmt"

func main() {

	fmt.Println("hey hi")
	random()
}

func random() {

  defer fmt.Println("zero")
  fmt.Println("one")
  fmt.Println("two")
  fmt.Println("three")


}