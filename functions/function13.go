package main 

import "fmt" 

func main() {
	g:= greeter {
		greeting : "hello",
		name: "go",
	}
	g.greet()
	fmt.Println( "the new name :", g.name)
}

type greeter struct {

	greeting string 
	name string
}
func (g *greeter) greet() {

	fmt.Println( g.greeting, g.name)

	g.name = "world"
}