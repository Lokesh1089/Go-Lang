// find string index

package main
import (
	"fmt"
	"regexp"
)

func main() {

	exp := regexp.MustCompile("van")
	fmt.Println(exp.FindStringIndex("vanheusen"))
	fmt.Println(exp.FindStringIndex("jockey"))
	fmt.Println(exp.FindStringIndex("panchavanam"))

}