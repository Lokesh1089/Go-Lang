package main 
import "fmt"
import "reflect"

type celsius float64
type farenheit float64

type defaultunit = farenheit


func main() {

	var f farenheit
	//var c celsius
    //c = 32
	f = 100
   // t =c
	t := f 
	fmt.Println(reflect.TypeOf(t))

} 