package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/brandoniu/lets-go/api"
	"github.com/brandoniu/lets-go/store"

	_ "github.com/lib/pq"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to communicate with database:", err)
	}

	store := store.New(db)

	app := api.NewAPI(store)

	app.Engine.Run(":8080")

	
}