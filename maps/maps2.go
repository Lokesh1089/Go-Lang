

package main

import "fmt"

func main(){

       departmentContactDetails := map[string] int {

		"mechanical": 212247, 
		"civil" : 214562,
		"ece" : 213645,
		"eee" : 214364,
		"cs" : 241354,
	  }

	  fmt.Println(departmentContactDetails)
	 
	  m := map[[5]int]string{}
	  k := make(map[string]int,10)

	 fmt.Println(m)
	 fmt.Println(k)

}