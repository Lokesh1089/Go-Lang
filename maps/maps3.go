// manipulation
	
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
	  departmentContactDetails["it"] = 214879
	  fmt.Println(departmentContactDetails)

	  departmentContactDetails["agri"] = 236589
	  departmentContactDetails["chemical"] = 245698

	  fmt.Println(departmentContactDetails)
	  


}