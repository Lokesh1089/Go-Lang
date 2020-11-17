package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
  
	read, err := ioutil.ReadFile("greet.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println( string(read))


	file2 , err := os.Open("test.txt")
	defer file2.Close()

	if err != nil {
		fmt.Println( err)
	}

	read2 := bufio.NewReader(file2) 

	r , err := read2.Peek(11)
	if err != nil {
		fmt.Println(err)
}

fmt.Println(string(r))

}