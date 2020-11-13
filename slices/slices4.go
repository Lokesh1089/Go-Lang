package main

import "fmt"

func main() {

	days := [] string {"sun", "mon", " tues", "wed", "thurs", "fri", " sat"}
	
	fmt.Println(days)

	weekdays := days[1:6]
	
	 days[2] = "tuesday"

	fmt.Println(weekdays)
	fmt.Println(days)
	 
}