// complex

package main
import ("fmt")

func main(){

	var a complex64 = 4i 
	fmt.Printf("%v, %T \n", a, a)

	
	var z complex64 = complex(7,3)
	fmt.Printf("%v, %T \n", z,z)

}