package main 
import "fmt" 
import "time"

func main() {
	var msg string = "vanakkam bro"
	
	go func() {
		fmt.Println(msg)
	}()
 
	msg = "hello"
time.Sleep(100*time.Millisecond)

}