package main
import "fmt"

//Product struct
type  Product struct {
	productID int 
	Name  string 
	Price int
}

func(p Product) print(){
	fmt.Println(p.productID, p.Name, p.Price)
}

func main(){

	pro1 := Product{productID:1 , Name: "Coconut Hair Oil " , Price: 180 }
	pro2 := Product{productID:2 , Name: "Kissan jam" , Price: 45 }
	pro1.print()
	pro2.print()

} 