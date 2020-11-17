	
	// peek read
	
	package main

	import (
		"fmt"
		"os"
		"bufio"
		
	)
	
	func main() {
	
	f, err := os.Open("greet.txt")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(f)

	b1 , err := reader.Peek(15)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println( string(b1))

}