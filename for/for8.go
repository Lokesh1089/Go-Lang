package main 
import "fmt"

func main(){

   stuMark := map[string]int {
					
				"kavi": 421,
				"shankar": 345,
				"murali": 456,
				"karthick": 199,

   }
   fmt.Println(stuMark)
   for k := range stuMark { fmt.Println(k)}
   for _,v := range stuMark { fmt.Println(v)}

}