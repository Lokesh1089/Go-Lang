package main 
import "fmt"

func main() {

	teachers := [...]string{ "mullai", "venba","kishore", "malar", "parthiban"}
	
	fmt.Println(teachers)
	fmt.Printf("teacher#3: %v\n", teachers[3])
	fmt.Printf("number of teachers: %v\n", len(teachers))
	

}