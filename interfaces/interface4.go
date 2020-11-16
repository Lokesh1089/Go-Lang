package main
import "fmt" 

type employee interface {

	expenses() int
}

type comstaff struct {

	salary, pf int 
}

type contractstaff struct {

	salary int
}

func ( c comstaff) expenses() int {

	return c.salary + c.pf
}

func ( c2 contractstaff) expenses() int {

	return c2.salary
}

func getExpenses(e employee) int {

	return e.expenses()
}

func main() {
	
	   c := comstaff{ salary: 18000, pf: 2500}
	   c2 := contractstaff {salary : 15000}

	   fmt.Printf("salary for company staff: %v Rs\n", getExpenses(c) )
	   fmt.Printf("salary for contract staff: %v Rs\n", getExpenses(c2) )
}