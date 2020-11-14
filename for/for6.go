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

   for k,v := range stuMark { fmt.Println(k,v)}

}