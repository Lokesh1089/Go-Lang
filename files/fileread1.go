package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"buffio"
	
)

func main() {

	data, err := ioutil.ReadFile("test.txt") 

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
	
}