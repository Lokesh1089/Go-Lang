package main 
import "fmt"

func main() {

  booksAndPrice := map[string] int {

	"javascript":  389,
	"learn c": 350,
	"learn go in 1 day" : 457,
 }
fmt.Println(booksAndPrice)

booksAndPrice["my sql"] = 399
fmt.Println(booksAndPrice)

fmt.Println(booksAndPrice["learn c"])
fmt.Println(booksAndPrice["javascript"])

}