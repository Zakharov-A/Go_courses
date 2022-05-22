package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1985@tcp(127.0.0.1:3306)/golang1") //3306
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Data installation

	insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Alex', 25)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	fmt.Println("Connected to MySQL")
}
