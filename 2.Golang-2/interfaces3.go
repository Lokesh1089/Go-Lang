package main 
import "fmt"
type printer interface{
	print() string
}
type scanner interface{
	scan() string
}
type scannerPrinter interface{
	printer
	scanner
}
type myPrinter struct {

}
type secondPrinter struct{

}
func(m myPrinter)  print() string{
	return "Printing first page..."
}
func(m myPrinter) scan()string {
	return "Scanning first page..."
}

func(s secondPrinter) print() string{
	return "Printing Five Pages"
}
func(s secondPrinter) scan() string{
	return "Scanning Five Pages"
}
func process(equip scannerPrinter){
	fmt.Println("Running Printer...", equip.print() )
	fmt.Println("Running Scanner...", equip.scan() )	
}
func main(){
	printer := myPrinter{}
	printer2 := secondPrinter{}
	process(printer)
	process(printer2)
}
