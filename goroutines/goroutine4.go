package main 
import "fmt" 
import "sync"

var wg = sync.WaitGroup {}

func main() {
	var msg string = "vanakkam bro"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	 msg = "hello"
	 wg.Wait()


}