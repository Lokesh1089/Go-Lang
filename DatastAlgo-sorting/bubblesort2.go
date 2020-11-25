package main
import (
	"fmt"
)

func main(){
	a:= []int{5,3,8,4,1,7,2,6}
    bubblesort(a)
}

func bubblesort(b []int){
 lengOfArray := len(b)

 for i:=0;i<lengOfArray-1;i++{
	 for j:=0;j<lengOfArray-i-1;j++{
		 if b[j]>b[j+1] {
		b[j], b[j+1] = b[j+1],b[j]
		 }
	 }
 }

 for _,v:= range b{

    fmt.Println("[",v,"]")

 }

}