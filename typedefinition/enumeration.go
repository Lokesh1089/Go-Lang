package main 

import "fmt"

// no of days in a week

type DayOfWeek int 
// days

const (

	sunday DayOfWeek = iota    //= 1 << iota
	monday
	tuesday 
	wednesday
	_
	anotherday
	thursday
	friday 
	saturday 
)

func main () {
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println( wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println( anotherday)
}
