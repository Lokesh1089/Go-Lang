
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
 
		for i:=1 ; i<10; i++{

			che<-i
		}

		close(che)
		waitgr.Done()
	}(che)
  waitgr.Wait()
}