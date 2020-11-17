package main 
import "fmt" 

func main() {

	var anfun func (int)
	anfun = func (number int) {

		if number == 0 {
			return
		} else {
			fmt.Println(number)

			anfun(number - 1)
		}
	}
	anfun(10)
}