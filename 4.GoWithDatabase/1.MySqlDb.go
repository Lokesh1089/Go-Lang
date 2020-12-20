package main
import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type Students struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Major string `json:"major"`
}

func main() {
	db, err := sql.Open("mysql", "root:2122@tcp(127.0.0.1:3306)/jarvis")
	if err != nil {
		fmt.Println("Cant connect to database")
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully Connected to Database")

	insert , err := db.Query("INSERT INTO students(name , major) VALUES ('Dharsha','Biology')")
	insert , err = db.Query("INSERT INTO students(name , major) VALUES ('Naveen','History')")
	
	if err != nil {
		fmt.Println("Cant Insert Value in the table")
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Inserted Value successfully")


	results, err := db.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println("Cant get data from database")
		panic(err.Error())
	}
	for results.Next(){
		var stu Students
		err = results.Scan(&stu.ID, &stu.Name, &stu.Major)
		if err != nil{
			panic(err.Error())
		}
		fmt.Println(stu)
	}

	results, err = db.Query("SELECT * FROM students WHERE id = 3")
	if err != nil {
		fmt.Println("Cant get data from database")
		panic(err.Error())
	}
	for results.Next(){
		var stus Students
		err = results.Scan(&stus.ID, &stus.Name, &stus.Major)
		if err != nil{
			panic(err.Error())
		}
		fmt.Println(stus)
	}


}