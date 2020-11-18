package main
import (
	"fmt"
	"regexp"
)

func main() {

	var content string= "once upon a time. one day we will win. ONE plus oNe is equal to two. try to eat one egg Once in a day"

	re := regexp.MustCompile("(?i)one(es)?")
	found := re.FindAllString(content, -1)
	fmt.Printf("%q\n", found)

	if found == nil{
		fmt.Println("no match found")

	}

	for _, word := range found {

		fmt.Printf("%s \n", word)
	}
}