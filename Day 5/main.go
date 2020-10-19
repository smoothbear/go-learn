package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	app := App{}
	app.Initialize(
		"127.0.0.0",
		"godb",
		"movies",
		)
	app.Run("8000")
	dbConnect()
}

func dbConnect() {
	db, err := sql.Open("mysql", "root:pwd@tcp (127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name FROM test WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
}
