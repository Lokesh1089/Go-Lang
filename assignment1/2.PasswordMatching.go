//password matching using regular expression 


package main
import (
	"fmt"
	"regexp"
)

func main(){
	fmt.Println(check("tamilrockers123"))
	fmt.Println(check("Tamilrockers&"))
	fmt.Println(check("#Tamilrockers123")) // valid password
	fmt.Println(check("Tamil*rockers"))
	fmt.Println(check("@Tamil23"))
}

func check( pw string) string{
   
	if len(pw)<9 {
		return "length of password is less than 9"
	}

	numbersvalidation, _ := regexp.MatchString(`[0-9]{1}`, pw)
	lowercasevalidation, _ := regexp.MatchString(`[a-z]{1}`, pw)
	symbolsvalidation,_ := regexp.MatchString(`[~!@#$%^&*?<>"}+_)("]{1}`, pw)
	uppercasevalidattion, _ := regexp.MatchString(`[A-Z]{1}`, pw)

	if numbersvalidation != true{
		return "add any number"
	}
	if lowercasevalidation != true {

		return "add any lowercase alphabet"
	}
	if uppercasevalidattion != true {

		return "add any uppercase alphabet"
	}
	if symbolsvalidation != true {

		return "add atleast any one of the symbols"
	}

	return "Given Password is valid"

}