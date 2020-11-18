
// find string submatch

package main
import (
	"fmt"
	"regexp"
)

func main() {

 str := regexp.MustCompile("p([a-z]+)ch")

 fmt.Println(str.FindStringSubmatch("peach") )
 fmt.Println(str.FindStringSubmatch("winndow"))
 fmt.Println(str.FindStringSubmatch("peach pouch"))
 fmt.Println(str.FindStringSubmatch("pouch") )


}