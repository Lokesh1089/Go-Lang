package main
import "fmt"

func main () {

	 names := map[int] string {

		1 : "radha",
		2 : "velmurugan",
		3 : "vetriselvan",
        4 : "raja" ,  
	}
	fmt.Println( names)

	leaders := names

	delete(leaders , 3)

	fmt.Println(leaders)
	fmt.Println(names)

	delete(leaders, 4)

	fmt.Println(leaders)
	fmt.Println(names)
}