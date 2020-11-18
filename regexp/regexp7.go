package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	cars := [...]string{"hyundai" , "mahindra" , "nissan", "toyota", "audi", "kia",}

	for _,brand := range cars{

		found, err := regexp.MatchString(".san", brand)
		if err!= nil {
			log.Fatal(err)
		}
		if found {
			fmt.Printf("%s matches\n", brand)
		} else{
			fmt.Printf("%s didn't match\n", brand)
		}
	}
}