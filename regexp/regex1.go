package main 

import (
"fmt"
"regexp"
)

func main() {
  
	re:= regexp.MustCompile("at$")

	fmt.Println(re.FindString("cat"))
	fmt.Println(re.FindString("antman"))
	 
}