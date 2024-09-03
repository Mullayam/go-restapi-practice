package main

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func (app *application) connectToDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", app.DSN)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to database Postgres!")
	return db, nil
}
