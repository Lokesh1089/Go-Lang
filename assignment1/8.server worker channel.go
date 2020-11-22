package main

import (
	"fmt"
	"time"
)

func main(){


	workers := make(chan string)
	jobs := make(chan int)

	emp := [5]string{"raj","mani","kannan","kesav","santhosh"}

	
   go func(){

		 for _,v := range emp {

			workers <- v
	
		 }
	defer close(workers)
	}()

	go	func(){

		for i:=1; i<6;i++{

			jobs <- i
			}
	defer close(jobs)
	}()

	for i:=1;i<6;i++{

		e := <-workers
		j := <-jobs


		fmt.Printf("employee < %s > started working with the project: %d\n", e,j)
		time.Sleep(time.Second)
		fmt.Printf("employee # %s # done the project: %d\n", e, j )
		time.Sleep(time.Second)
	}
   

}