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
	db, err := sql.Open("mysql", "root:9999@tcp(127.0.0.1:3306)/testdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	fmt.Println("testing db....")

	rows, err := db.Query(`Select * from person where where Pid=26;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		timeline := Timeline{}
		err = rows.Scan(&timeline.Id, &timeline.Content)
		if err != nil {
			panic(err)
		}
		fmt.Println(timeline)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
