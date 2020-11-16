package main 

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0 
var m = sync.RWMutex{}
func main () {
	runtime.GOMAXPROCS(100)
	for i:= 0 ; i<10 ; i++{
		wg.Add(2)
		m.RLock()
		go sayhello()
		m.Lock()
		go increament()
	}
	wg.Wait()
}

func sayhello(){
	
	fmt.Printf("hello #%v \n", counter)
	m.RUnlock()
	wg.Done()
}

func increament() {
	
	counter++
	m.Unlock()
	wg.Done()
}