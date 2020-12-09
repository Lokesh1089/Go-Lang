package main
import "fmt"

type furniture interface{
	get() 
}
type sofa struct {
	numOfSeat int
	price int
	manufacturer string 
}
type diningTable struct {
	numOfSeat int
	price int
	manufacturer  string
}
func(s sofa) get(){
	fmt.Println(s.numOfSeat, s.price, s.manufacturer)
}
func(d diningTable) get(){
	fmt.Println(d.numOfSeat,d.price,d.manufacturer)
}
func main() {
	s := sofa{numOfSeat: 4 , price : 9599, manufacturer : "Chozhan Wood Works" }
	dt := diningTable{numOfSeat: 8, price:35000, manufacturer : "Vetri Wood Works Pvt Ltd"}

	fur := [] furniture{&s,&dt}
	
	for _,v := range fur {
		v.get()
	}


}