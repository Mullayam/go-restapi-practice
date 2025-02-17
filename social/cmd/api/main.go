package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mullayam/social/internal/db"
	"github.com/mullayam/social/internal/env"
	"github.com/mullayam/social/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Panicln("Loaded Environment Variables")

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:    env.GetString("DB_ADDR", "postgres://test:testlocalhost:5432/social?sslmode=disable"),
			maxOpen: env.GetInt("DB_MAX_OPEN", 25),
			maxIdle: env.GetInt("DB_MAX_IDLE", 25),
		},
	}
	db, err := db.New(cfg.db.addr, cfg.db.maxOpen, cfg.db.maxIdle, "5m")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Panicln("Connected to DB")
	store := store.NewPostgresStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
