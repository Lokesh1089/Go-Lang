// replace all string 


package main
import (
	"fmt"
	"regexp"
)

func main() {

	 re := regexp.MustCompile("ise")

	 s:= "ize"

	 fmt.Println(re.ReplaceAllString("realise", s))
	 fmt.Println(re.ReplaceAllString("organise", s))
	fmt.Println(re.ReplaceAllString("analyse", s))	
}