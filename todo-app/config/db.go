package config

import (
	"database/sql"
	"os"
)

type Repository struct {
	DB *sql.DB
}

func NewConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
