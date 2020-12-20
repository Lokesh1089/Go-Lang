package main
import (
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)
type Departments struct{
	Deptid int `json:"deptid"`
	Dname string `json:"dname"`
	Dlocation string `json:"dlocation"`
}

func main() {
	connParams := "host=localhost port=5432 dbname=test user=postgres password=4095 sslmode=disable"
	db , err := sql.Open("postgres", connParams)
	if err!= nil {
		panic(err.Error())
		}

	defer db.Close()	
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected To Database")
	
	ins , err := db.Query("INSERT INTO departments(deptid, dname , dlocation ) VALUES(1,'Production','Oragadam')")
	ins , err = db.Query("INSERT INTO departments(deptid, dname , dlocation ) VALUES(2,'Quality','Sriperumbadur')")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Inserted Departments Data Successfully")
	defer ins.Close()


	result , err := db.Query("SELECT * FROM departments ;")
	if err != nil {
		panic(err.Error())
	}
	for result.Next(){
		var dept Departments
		err := result.Scan(&dept.Deptid, &dept.Dname, &dept.Dlocation)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(dept)
	}


}