// slicing weekdays and weekend from weekss


package main 

import "fmt"

func main() {

	days := [] string { "sun", "mon", "tues", "wed", "thurs", "fri", "sat"}

	fmt.Println(days)

	weekdays := days[1:6]
	fmt.Println(weekdays)

	//weekend := append(days[:1], days[6:]...)
	//fmt.Println(weekend)

	//fmt.Println(days)

	weekend1 := days[:1]
	weekend2 :=	days[6:]
	 
	fmt.Println(weekend1,  weekend2)

	fmt.Println(days)

}