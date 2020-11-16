package main 
 import "fmt"
 

 type shapes interface {

     area() float64
 }

 type square struct {

	side float64
 }
 type triangle struct {

	base , height float64
 }

 func ( s square) area() float64 {

	return s.side * s.side
 }

 func ( t triangle) area() float64 {

	return 0.5*t.base*t.height
 }

 func getArea(s shapes) float64 {

	return s.area()
 }

 func main () {

	s := square { side : 25}
	t := triangle {base:10 , height: 8}
	fmt.Println("area of sqaure:", getArea(s))
	fmt.Println("area of triangle :", getArea(t))
 }