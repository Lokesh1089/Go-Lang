

package main

import (
	"fmt"
	"time"
)


func main() {

	 totalProjects :=7
	emp := make(chan int, totalProjects)
	status := make(chan int, totalProjects)

	for a := 1; a <=6; a++ {

		go employee(a, emp, status)
	}

	for b := 1; b <= totalProjects; b++ {
		emp <- b
	}
	for c := 1; c <= totalProjects ; c++ {
		<-status
	}

	close(emp)
}

func employee(no int, emp <-chan int, status chan<- int) {
	for a := range emp {
		fmt.Println("employee",no, "is working with project no", a)
		time.Sleep(time.Second*3)
		fmt.Println("employee", no, "has done the project", a)
		status <- 1
	}
}