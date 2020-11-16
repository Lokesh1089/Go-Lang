package main 
import "fmt"
import "math" 

type shape interface {
	 area() float64
}

type circle struct {

	 radius float64
}

type rectangle struct {

	height , width  float64
}

func (r rectangle) area() float64 {

	return r.height * r.width

}

func (c circle) area() float64 {

	return math.Pi * c.radius*c.radius
}
func getArea(s shape) float64{

	return s.area()
}

func main() {
	 
	c := circle{radius: 6.0}
	r := rectangle{height :12.0,width : 8.0}

	fmt.Printf("circle area: %f\n", getArea(c))
	fmt.Printf("rectangle area: %v\n", getArea(r))
}