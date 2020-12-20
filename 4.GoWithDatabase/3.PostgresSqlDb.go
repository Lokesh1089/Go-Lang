package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
type Employee struct {
	EmpID int `json:"empid"`
	Ename string `json:"ename"`
	Dept string `json:"dept"`
}
func main() {
	conn := "host=localhost port=5432 dbname=test user=postgres password=4095 sslmode=disable"
	db , err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to database")

	insert , err := db.Query("INSERT INTO employee VALUES(1,'Ajay', 'Production');")
	insert , err = db.Query("INSERT INTO employee VALUES(1,'Bharathi', 'Quality');")

	if err != nil {
		panic(err.Error())	
	}

	defer insert.Close()
	fmt.Println("Inserted Successfully")


	get , err := db.Query("SELECT * FROM employee;")
	if err != nil{
		panic(err.Error())
	}
	for get.Next(){
		var emp Employee
		err := get.Scan(&emp.EmpID, &emp.Ename, &emp.Dept)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(emp)
	}



}