package main
import "fmt"
//Human interfaces
type Human interface{
	getInfo() 
	setInfo()
}
//Student struct
type Student struct {
	rollNo int
	name string
	total float32
}
//Employee struct
type Employee struct {
	id int
	name string
	salary float32
}
func(s *Student) setInfo(){
	fmt.Scan(&s.rollNo,&s.name,&s.total)
}
func(e *Employee) setInfo(){
	fmt.Scan(&e.id,&e.name,&e.salary)
} 
func(s *Student) getInfo(){
	fmt.Println(s.rollNo,s.name,s.total)
}
func(e *Employee) getInfo(){
	fmt.Println(e.id,e.name,e.salary)
}
//interface as parameter
func process(h Human){
	h.setInfo()
	h.getInfo()
}

func main(){

	// stu :=Student{ 1, "Parthiban", 485 }
	// emp :=Employee{101, "Vetrivel", 20000 }
	// stu.getInfo()
	// emp.getInfo()
	// humans := []Human{&stu, &emp}
humans := []Human{new(Student),new(Employee) }

	for _,v := range humans {
		process(v)
	}

}