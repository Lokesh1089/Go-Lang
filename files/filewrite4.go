package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	 
	greetings , err := os.Create("greet.txt")
	defer greetings.Close() 
	if err != nil{
		fmt.Println(err)
	}
	
	greetings.WriteString("good morning !")

	greetings.Sync()

	person := bufio.NewWriter(greetings)

	person.WriteString(" manivannan")

	person.Flush()


}