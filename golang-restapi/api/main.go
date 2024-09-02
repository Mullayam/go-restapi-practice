package main

import (
	"fmt"
	"log"
	"net/http"
)

type application struct {
	Domain string
}

const PORT = 8080

func main() {
	// set app config
	var app application
	app.Domain = "localhost"
	log.Println("Starting server on port", PORT)
	// set db config
	// set logger
	// set routes
	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
