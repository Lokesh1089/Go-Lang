package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type CarSales struct{

	Brand string
	Model string
	SalesRepId int64 
	SaleTime time.Time
}

func main() {

	car := CarSales{
		Brand: "hyundai",
		Model: "verna",
		SalesRepId: 247457,
		SaleTime: time.Now() ,
	}

	carjs, err := json.MarshalIndent(car,"", "  ")

	if err!= nil{
		fmt.Println(err)
	}

	a := string(carjs)
	
	fmt.Println(a)


}