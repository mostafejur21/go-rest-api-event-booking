package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// this will open or create a db. in this case it is sqlite 3
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}

func createTable() {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER
    )
    `

	result, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}
	log.Println(result)
}
