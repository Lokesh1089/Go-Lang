
// & 

package main
import "fmt"



func main(){

students := struct{ name string} { name:  " raj kumar"}
 fmt.Println( students)

 newstudents := students 
 newstudents.name = "donald triumph"

 fmt.Println(newstudents)
 fmt.Println( students)

 stu := &students
 stu.name = "joe biden"
 fmt.Println(stu)
 fmt.Println(students)

}