package main 
import "fmt"

type User struct{
	Name string 
	Email string
	Age int
}

func main(){

	ram := User{Name: "Ram Kumar", Email:"ram@xyx.com", Age: 23 }
	fmt.Println(ram)
	fmt.Printf("%+v \n", ram )

	mukund := new(User) 
	mukund.Name = "Abhinav Mukund"
	mukund.Email = "amuk123@xyz.com"
	mukund.Age = 35
	fmt.Println(mukund)

	kavin := &User{Name: "Kavin", Email: "kavin123@abc.com", Age: 26 }
	fmt.Println(kavin)
}