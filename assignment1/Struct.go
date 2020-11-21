package main
import (
	"fmt"
)

type MyCustomer struct{

	name string
	locaion string
	cusomerid int
	isRegularCustomer bool
}

func main() {
	 
	cus1 := MyCustomer{

		name:"Gokul",
		locaion: "Tirupattur",
		cusomerid: 47,
		isRegularCustomer: true,
	}

	cus2 := MyCustomer{
		name: "Vignesh",
		locaion: "krishagiri",
		cusomerid: 21,
		isRegularCustomer: false,
	}

	 var cus3 MyCustomer
	 cus3.name = "Prakash"
	 cus3.locaion = "uthangarai"
	 cus3.cusomerid = 17
	 cus3.isRegularCustomer= true

  
	fmt.Println(cus1)
	fmt.Println(cus2)
	fmt.Println(cus3)

	fmt.Println(cus2.locaion)
	fmt.Println(cus1.isRegularCustomer)
	fmt.Println(cus3.name)

}