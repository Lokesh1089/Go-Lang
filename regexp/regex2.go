// string index

package main 
import (
	"fmt" 
	"regexp"
)

func main () {
  
	str := regexp.MustCompile("tel")

	fmt.Println(str.FindStringIndex("telephone"))
	fmt.Println(str.FindStringIndex("market"))
	fmt.Println(str.FindStringIndex("airtel"))

}