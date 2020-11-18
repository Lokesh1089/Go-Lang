package main
import "fmt" 

func main() {

	var x string = "good morning"
	 var b []byte

	 b = []byte(x)
	  fmt.Println(b)

	 var z string = string(b) 
     fmt.Println(z)

}