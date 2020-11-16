package main 
import(
	 "fmt"
	 "sync"
)	
 
var waitg = sync.WaitGroup{}

func main() {

	mychan := make(chan int)
	for j:= 0 ; j<5; j++ {
	waitg.Add(2)
		go func() {
			k := <- mychan
			fmt.Println(k)
			waitg.Done()
		}()
		go func() {
			mychan <-33
			waitg.Done()
		}()	
	}
	waitg.Wait()
}