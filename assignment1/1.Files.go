package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	//	"os"
)

func main() {

	read1 , err := ioutil.ReadFile("colours.txt")  // file reading ioutil
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(string(read1))
//-------------------------------------------------------------------------------
	file1 , err := os.Open("quotes.txt")   // file reading by bufio 

	if err != nil {
		fmt.Println(err)
	}

	read2 := bufio.NewReader(file1)

	bytRead , err := read2.Peek(10)  // peek and read
	if err!=nil{
		fmt.Println(err)
	}
	
	fmt.Println(string(bytRead))


  // creating file ----------------------------------------------------------------------

	createFile1, err := os.Create("numbers.txt")
	defer createFile1.Close()
    if err != nil{
		fmt.Println(err)
	}

	createFile1.WriteString("one two three four five six seven eight nine ten......  ")

	createFile1.Sync()

//  bufio new writer 

	endStatement := bufio.NewWriter(createFile1)
	
	endStatement.WriteString("counting is over lets rock guys !")

	endStatement.Flush()

}




