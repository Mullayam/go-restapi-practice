package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type application struct {
	Domain string
	DSN    string
	DB     *sql.DB
}

const PORT = 8080

func main() {
	// set app config
	var app application
	app.Domain = "localhost"
	log.Println("Starting server on port", PORT)

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", "Postgres connection string")
	flag.Parse()

	// set db config
	db, e := app.connectToDB()
	if e != nil {
		log.Print(e)
		log.Fatal(e)
	}
	app.DB = db
	defer db.Close()

	// set logger
	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
