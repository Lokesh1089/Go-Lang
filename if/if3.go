package main

import( 
	"fmt"
	"math"
)

func main() {

mynum := 0.1
if mynum == math.Pow(math.Sqrt(mynum),2){
	fmt.Println("these are same")
} else{

	fmt.Println("these are different")
}

mynum1 := 0.123
if mynum1 == math.Pow(math.Sqrt(mynum1),2) {
	fmt.Println("these are same")
} else{

	fmt.Println("these are different")
}

}