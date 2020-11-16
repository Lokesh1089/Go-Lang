package main 
import "fmt"
import "sync"

var wg = sync.WaitGroup{}

func main () {
  
	ch := make(chan int)
	wg.Add(2)
	go func(){
		i := <- ch
		fmt.Println(i)
		wg.Done()
		
	}()	
	go func(){
		 k := 33
		ch <- k
		 k = 11

		wg.Done()
	}()	

	wg.Wait()

} 