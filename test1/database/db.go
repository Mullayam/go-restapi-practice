package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitConnection() {
	DB, err := sql.Open("sqlite3", "./localDatabase.db")
	if err != nil {
		panic("failed to connect database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	createTables()
	// return DB, nil
}
func createTables() {
	createEventsTable := `
			CREATE TABLE IF NOT EXISTS events (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(255) NOT NULL,
				description VARCHAR(255) NOT NULL,
				location VARCHAR(255) NOT NULL,
				date_time DATETIME NOT NULL
				user_id INTEGER
			);	
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("failed to create table")
	}

}
