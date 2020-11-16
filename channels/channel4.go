//sent only recive only

package main 
import(
 "fmt"
 "sync"
)

var waitgr = sync.WaitGroup{}

func main() {
  
	che := make(chan int)

	waitgr.Add(2) 
	go func ( che <-chan int){
		i := <- che
		fmt.Println(i) 
		waitgr.Done()
	} (che)

	go func (che chan<- int) {
 
		che <- 123
		waitgr.Done()
	}(che)
  waitgr.Wait()
}