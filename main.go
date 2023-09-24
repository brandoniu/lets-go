package main

import (
	"database/sql"
	"log"

	"github.com/brandoniu/lets-go/api"
	"github.com/brandoniu/lets-go/store"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=books sslmode=disable password=mypassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	store := store.New(db)

	app := api.NewAPI(store)

	app.Engine.Run(":8080")

	
}