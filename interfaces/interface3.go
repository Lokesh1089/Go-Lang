package main 
import "fmt" 

type vehicle interface {

	distance() float64
}

type bike struct {
	speed, time  float64

}

type car struct {
	speed , time float64
}

func ( b bike) distance() float64 {

	return b.speed*b.time
}
func( c car) distance () float64{

	return c.speed*(c.time)
}

func getSpeed(v vehicle) float64 {

	return v.distance()
}

func main() {

	b := bike{speed: 110 , time : 2}
	c:= car{speed: 85 , time: 3}
	fmt.Printf("distance covered by bike %vkms \n", getSpeed(b))
	fmt.Printf("distance covered by car %vkms \n", getSpeed(c))
}