// receiving more than one data from sender


package main 
import(
 "fmt"
 "sync"
)

var waitgr = sync.WaitGroup{}

func main() {
  
	che := make(chan int, 50)

	waitgr.Add(2) 
	go func ( che <-chan int){
		i := <- che
		fmt.Println(i) 
		i = <- che
		fmt.Println(i) 
		i = <- che
		fmt.Println(i) 
		waitgr.Done()
	} (che)

	go func (che chan<- int) {
 
		che <- 123
		che <- 55
		che <- 33
		waitgr.Done()
	}(che)
  waitgr.Wait()
}