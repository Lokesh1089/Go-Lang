
// complex

package main 
import "fmt"

func main(){
   var v complex64 = 12+4i 
   const k complex64 = 4 + 3i 

   fmt.Println(v+k)
   fmt.Println(v-k)
   fmt.Println(v*k)
   fmt.Println(v/k)
   //fmt.Println(v%k)

}