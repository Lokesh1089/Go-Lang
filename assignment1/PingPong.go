// ping pong 

package main
import (
	"fmt"
	"time"
)	

func pinger( pingchan <-chan int, pongchan chan<- int ){
	
  for{
	a := <- pingchan
	fmt.Println("ping", a)           //ping got the ball 
	time.Sleep(time.Second)
	pongchan <- 1				// ping hits the ball to pong
  }
}

func ponger( pingchan chan<- int , pongchan <-chan int ){

	for {
	 b := <- pongchan
	 fmt.Println("pong", b)   // its a pongs turn 
	 time.Sleep(time.Second)
	 pingchan <- 0            // hitting hard the ball to ping
	}
}	 
func main(){

	ch1 :=  make(chan int)
	ch2 := make(chan int)
 
	go pinger(ch1, ch2)
	go ponger(ch1,ch2)
 
	ch1 <-0  //game starts here
	for {
		time.Sleep(time.Second)  // for tiredless play
	}
	
 }