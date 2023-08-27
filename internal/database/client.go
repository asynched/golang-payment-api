package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB = nil

func CreateClient() *sql.DB {
	if database != nil {
		return database
	}

	db, err := sql.Open("sqlite3", "./database.sqlite")

	if err != nil {
		panic("Could not establish a connection with the database")
	}

	database = db

	return db
}
