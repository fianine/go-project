package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Laruku23!@tcp(localhost:3306)/go_project")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
