package main 
import "fmt" 

type Celcius float32

type Farenheit float32

func (c Celcius) ToFarenheit() Farenheit {

	return Farenheit(9.0/5.0*c + 32)
}
func main() {


	var  c Celcius
	c= 38.55
	 f := c.ToFarenheit()
	 fmt.Println(f)
} 