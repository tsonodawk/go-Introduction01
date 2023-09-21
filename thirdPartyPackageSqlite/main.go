package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// cmd := `CREATE TABLE IF NOT EXISTS persons(
	// 	name STRING,
	// 	age INT)`

	// _, err := Db.Exec(cmd)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err = Db.Exec(cmd, "hanako", 19)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	cmd := "SELECT * FROM persons where age = ?"

	row := Db.QueryRow(cmd, 19)
	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			println("No row")
		} else {
			println(err)
		}
	}

	fmt.Println(p.Name, p.Age)
}
