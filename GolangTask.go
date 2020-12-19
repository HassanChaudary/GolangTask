package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Records struct {
	PID   string
	Name  string
	email string
	phone string
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, records := range records {
		data := Records{
			PID:   records[0],
			Name:  records[1],
			email: records[2],
			phone: records[3],
		}
		//fmt.Println(data.PID + " " + data.Name + " " + data.email + " " + data.phone)

	}

	db, err := sql.Open("mysql", "root:9999@tcp(127.0.0.1:3306)/testdb")

	file, err := os.Open("C://Users/Hassan/Desktop/people.csv")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("INSERT INTO users (PID,Name,email phone) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		records := Records{}
		err = rows.Scan(&records.PID, &records.Name, &records.email, &records.phone)
		if err != nil {
			panic(err)
		}
		fmt.Println(records)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
