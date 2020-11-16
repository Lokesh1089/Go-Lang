package main 

import (
	"fmt"

	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0 

func main () {
	for i:= 0 ; i<10 ; i++{
		wg.Add(2)
		go sayhello()
		go increament()
	}
	wg.Wait()
}

func sayhello(){
	fmt.Printf("hello #%v \n", counter)
	wg.Done()
}

func increament() {
	counter++
	wg.Done()
}