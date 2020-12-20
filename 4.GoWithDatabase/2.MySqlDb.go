package main 
import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	)

type Books struct{
	BookID int `json:"bookid"`
	Title string `json:"title"`
	Author string `json:"author"`
	}

func main() {

	db , err := sql.Open("mysql", "root:2122@tcp(127.0.0.1:3306)/jarvis")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(" Goood Job ! Successfully connected to database")

	insert , err := db.Query("INSERT INTO books(bookid,title,author) VALUES(1,'Kadal Pura', 'Sandilyan')")
	insert , err = db.Query("INSERT INTO books(bookid,title,author) VALUES(2,'Ponniyin Selevan', 'kalki')")
	insert , err = db.Query("INSERT INTO books(bookid,title,author) VALUES(3,'Gopalla Gramam', 'Raja Narayanan')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	get , err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err.Error())
	}
	for get.Next(){
		var boo Books
		err := get.Scan(&boo.BookID, &boo.Title, &boo.Author)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(boo)

	}
	
	getOne , err := db.Query("SELECT * FROM books WHERE author= 'kalki'")
	if err != nil {
		panic(err.Error())
	}
	for getOne.Next(){
		var bOne Books
		err := getOne.Scan(&bOne.BookID, &bOne.Title,&bOne.Author)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(bOne)
	}

}	