package main
import "fmt"
import "sync"

var vg = sync.WaitGroup{}

func main () {
 
   che := make(chan string)
   vg.Add(2)

   go func () {
	  
	  greeting := <- che
	  fmt.Println(greeting)
	  vg.Done()
	}()

	go func() {
		 che <- "good morning bro!"
		vg.Done()

	}()
	vg.Wait()

}