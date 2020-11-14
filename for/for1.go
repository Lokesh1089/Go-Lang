package main
import "fmt"


func main(){
	
	for i:= 0; i<5; i++{

		fmt.Println(i)
	}

	for k,j := 0,0 ; j < 8 ; j,k = j+2, k+4{

		fmt.Println(j,k)

	}

}