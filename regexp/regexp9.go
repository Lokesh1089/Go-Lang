
// split

package main
import ("fmt"
"log"
"regexp"
"strconv"
)

func main() {

  var nums = "11, 23, 45, 21, 12, 1, 3, 7, 37, 26"

   sum := 0
  rex := regexp.MustCompile(",\\s*")
	values := rex.Split(nums, -1)
	for _,numval := range values{
		n, err := strconv.Atoi(numval)

			sum += n

			if err != nil{
				log.Fatal(err)
			}
	}
  fmt.Println(sum)
}