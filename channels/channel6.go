package main 
import (
	"fmt" 
	"sync"
)

var wg = sync.WaitGroup{}
func main() {
ch := make(chan string, 50) 

wg.Add(2)
go func ( ch <-chan string) {
   greetings := <- ch
	fmt.Println(greetings)

	 	greetings = <- ch
	   fmt.Println(greetings)
	   
   	greetings = <- ch
   	fmt.Println(greetings)
   wg.Done()
}(ch)

go func ( ch chan<- string) {

	 ch<- "good morning" 
	 ch <- "good afternoon"
	ch <- "good evening"
	wg.Done()
}(ch)

wg.Wait()
}