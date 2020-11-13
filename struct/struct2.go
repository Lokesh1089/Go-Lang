package main

import "fmt"

type students struct {

	sName string
	sID int
	major string
}

func main () {

		details := students {

			sName: "Manirathinam",
			sID:  22,
			major: "mechanical",
		}

		details2 := students{

			sName:  "ashok kumar",
			sID:  41,
			major : "civil", 
		}
fmt.Println( details)
fmt.Println(details2)

fmt.Println( details.major)

}