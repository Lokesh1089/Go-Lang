package main

import (
	"fmt"
	"os"
)

func main () {

	f, err := os.Create("newtextfile.txt")
	defer f.Close()
    if err != nil {

		 fmt.Println(err)
	}
	f.WriteString("hey go this is me again")

	f.Sync()	
}