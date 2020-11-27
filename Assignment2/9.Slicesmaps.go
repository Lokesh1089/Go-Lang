package main

import (
	"fmt"
)


func main() {
fmt.Println("--------------------slice-----------------------------------")
	slices := [] int {9,8,7,6,5,4,3,2,1}
	fmt.Println("initial slices",slices)
	slices = append(slices, 0,-1,2)
	fmt.Println("After append",slices)
	slices = slices[:]
	b := slices[3:]
	fmt.Println("slicing view",b)
	c :=slices[8:11]
	fmt.Println("Slicing view",c)
	slices[0] = 10
	fmt.Println("After Updating value",slices)

fmt.Println("--------------------map-----------------------------------")

maps := map[string] int{
						"Ashwin": 1089,
						"Varun": 985,
						"krithika": 890,
						"Raj":1154,
						"kanmani": 1126,
						"Arun": 774,
						"Dinesh":1071,
						 }
						 
	fmt.Println(maps)
	maps["Elavarasan"] = 912
	fmt.Println("Adding new element",maps)
	delete(maps, "krithika")
	fmt.Println("After Delete",maps)
	checkList, ok := maps["Raj"]			//checking
	fmt.Println(checkList, ok)
	fmt.Println("Printing dinesh mark", maps["Dinesh"])
}