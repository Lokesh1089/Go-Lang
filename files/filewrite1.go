package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

func main() {
   
	 message := []byte(" hello go im writing message to you")
	 err := ioutil.WriteFile("new.txt", message, 0644) 

	 if err != nil {
		 fmt.Println(err)
	 }
}