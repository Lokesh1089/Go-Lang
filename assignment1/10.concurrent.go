package main

import (
	"fmt"
)

func main() {
 
	out1 := make(chan int)
	out2 := make(chan string)

	 go func(){
		for i:=1;i<6;i++{
		out1 <- i
			}
		close(out1)	
	}()

	go func(){
		for i:=1;i<6;i++ {
		out2 <- "yeah! im from channel two"

		}
	close(out2)
	}()

   for {
		num, isopen := <- out1
		if !isopen {
			break
		}
		fmt.Println(num)

	}

	for {
		msg, isopen := <- out2
		if !isopen {
			break
		}
		fmt.Println(msg)
	}

}
