package main 
import (
	"fmt"
	"runtime"
)
func main(){

	runtime.GOMAXPROCS(10)
 
  fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}


	
