package main

import (
	"fmt"

	"regexp"
)
func main() {

	websites := [...] string {"kloudlearn.com", "tngov.in", "tamilrockers.ws", "intime.org"}

	a := regexp.MustCompile("(\\w+)\\.(\\w+)")

	for _, v := range websites{

		parts := a.FindStringSubmatch(v)



		for k := range parts{
		fmt.Println(parts[k])
		}

		fmt.Println("##############################")
	}

}