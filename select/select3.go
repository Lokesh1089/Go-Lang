package main 
import ( 
	"fmt" 
	"time" 
)

func main() {

	ch1 := make(chan int) 
	ch2 := make( chan int) 
	ch3:= make( chan int) 

 go func () {

	ch1 <- 24
	time.Sleep(1* time.Second )

 }()
 
 go func () {

	ch2 <- 11
	time.Sleep(1* time.Second )

 }()

 go func () {

	ch3 <- 55
	time.Sleep(1* time.Second )

 }()
 for i:= 0; i<3 ; i++ {

	select {
	case value1 := <- ch1 :
		fmt.Println("value one :", value1)
	case value2 := <- ch2 :
		fmt.Println("value two :", value2)
	case value3 := <- ch3 :
		fmt.Println("value three :" , value3)		

	}
 }


}