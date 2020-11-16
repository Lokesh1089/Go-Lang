
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
		for {
			if i , ok := <- che ; ok {
				fmt.Println(i)
			} else{
				break
			}
		}
		
		waitgr.Done()
	} (che)

	go func (che chan<- int) {
 
		che <- 1
		che <- 2
		che <- 3
		che <- 4
		che <- 5
		che <- 6

		close(che)
		waitgr.Done()
	}(che)
  waitgr.Wait()
}