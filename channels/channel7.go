// using for to receive more values from sending channel


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
		for i := range che {
			fmt.Println(i) 
		}
		
		waitgr.Done()
	} (che)

	go func (che chan<- int) {
 
		che <- 123
		che <- 55
		che <- 22
		close(che)
		waitgr.Done()
	}(che)
  waitgr.Wait()
}