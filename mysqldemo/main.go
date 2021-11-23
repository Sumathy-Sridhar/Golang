package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goonn"dbc)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("CREATE TABLE dbdemo(id INT NOT NULL, name VARCHAR(20),PRIMARY KEY (ID));")
	_, err = db.Query("INSERT INTO dbdemo VALUES(003,'Mathew')")
	_, err = db.Query("INSERT INTO dbdemo VALUES(004,'Delvin')")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}
