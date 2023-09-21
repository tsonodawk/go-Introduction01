package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err := sql.Open("postgres", "user=postgres password=postgres dbname=testdb sslmode=disable")

	if err != nil {
		log.Panicln(err)
	}

	defer Db.Close()

	// cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	// _, err = Db.Exec(cmd, "hanako", 19)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	cmd := "SELECT * FROM persons where age = $1"

	row := Db.QueryRow(cmd, 19)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			println("No row")
		} else {
			println(err)
		}
	}

	fmt.Println(p.Name, p.Age)
}
