package main

import "fmt"

func main(){
switch i:= 2+2; i {
case 1,2,3: 
   fmt.Println("one , two , three")
   fmt.Printf("%v", i)

case 4,5,6:
fmt.Println(("four, five, six"))
fmt.Printf("%v", i)
default :
   fmt.Println("above seven")

}


}