// ecomposition relation


package main 
import "fmt"

type Animal struct{

	name string
	origin string
}

type Bird struct{

	Animal
	speedKPH float32
	canFly bool
}

func main(){

	b := Bird{}
	b.name = "emu"
	b.origin = "australia"
	b.canFly = false
	b.speedKPH = 48

	fmt.Println(b.name)


	 c := Bird{

			Animal : Animal{ name: "wood pecker", origin: "australia"},
			speedKPH:  80,
			canFly: true ,
			 }
	

			fmt.Printf(c.name)

	
}