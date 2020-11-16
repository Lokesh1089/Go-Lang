package main 
import "fmt"
import "math"

type point struct{

 x, y float64

}

func ( p1 point) DistanceFrom(p2 point) float64 {
	 return math.Hypot(p2.x- p1.x, p2.y - p1.y)
}
	
func(p3 *point) scale(factor float64) {

	p3.x = p3.x * factor
	p3.y = p3.y * factor

}

func main(){

	p1 := point{1,1}
	p2 := point{1,5}

	k := p1.DistanceFrom(p2)
	fmt.Println(k)

	p3 := point {2,4} 

	(&p3).scale(3)

	fmt.Println(p3)

}