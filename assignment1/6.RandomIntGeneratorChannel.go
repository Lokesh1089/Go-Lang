// random int generator using channel

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	ch1 := make(chan int)
  
	go intgen(ch1)
	ch1 <- 188
	time.Sleep(time.Second)

}

func intgen(num <-chan int){

	a := <- num
	fmt.Println("Random number is :", rand.Intn(a) )
	
	time.Sleep(time.Second)
}