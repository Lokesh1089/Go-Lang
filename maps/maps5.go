
// deleting value from map
package main

import "fmt"
func main() {

 students := map[string] int {
	 
	"kannan" : 487,
	"murali" : 398,
	"mithun" : 274,
	"karthik": 423,
	"arun" : 199,

 } 
  
 fmt.Println(students)

 delete(students, "mithun")
 delete(students , "kannan")
fmt.Println(students)

 delete(students, "arun")
 fmt.Println(students)

 }