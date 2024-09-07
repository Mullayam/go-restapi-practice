package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	dbrepo "github.com/test/api/db"
	"github.com/test/api/repository"
)

type application struct {
	Domain       string
	DSN          string
	DB           repository.DatabaseRepository
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

const PORT = 8080

func main() {
	// set app config
	var app application

	log.Println("Starting server on port", PORT)

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "enjoys.in", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "enjoys.in", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "localhost", "domain")

	flag.Parse()

	// set db config
	db, e := app.connectToDB()
	if e != nil {
		log.Print(e)
		log.Fatal(e)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: db}

	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Hour * 24,
		RefreshExpiry: time.Hour * 24 * 7, // 7 days
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}

	// set logger
	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
