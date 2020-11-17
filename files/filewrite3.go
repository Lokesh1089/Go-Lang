package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	file, err := os.Create("anothernewfile.txt")
	defer file.Close()
    if err != nil {

		 fmt.Println(err)
	}
   
	wr := bufio.NewWriter(file)
	wr.WriteString("hello ! how are bro")

	wr.Flush()
}