package main

import (
	"fmt"
	"sync"
	"time"
)

var waitgr = sync.WaitGroup{}

func main() {
  
	che := make(chan string)

	waitgr.Add(2) 
	go func ( che <-chan string){
		for {
			if i , ok := <- che ; ok {
				fmt.Println(i)
			} else{
				break
			}
		}
		
		waitgr.Done()
	} (che)

	go func (che chan<-string) {
 
		for{

			che<-"Hello"
			time.Sleep(time.Second)
			che<-"welcome to bommi's bakery"
			time.Sleep(time.Second)
			che<-"we have 18 variety of cakes"
			time.Sleep(time.Second)
			che<-"& 10 variety of puffs "
			time.Sleep(time.Second)
			che<-"Thank you visit  again"
			time.Sleep(time.Second*3)


			
		}

		close(che)
		waitgr.Done()
	}(che)
  waitgr.Wait()

  time.Sleep(time.Second)

}