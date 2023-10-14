package main

import (
	"log"
	"os"

	"github.com/brandoniu/lets-go/api"
	"github.com/brandoniu/lets-go/repository"
	"github.com/brandoniu/lets-go/service"
)

func main() {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL not set")
	}
	
	bookRepo, err := repository.NewPostgresBookRepository(connString)
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	BookService := service.NewBookService(bookRepo)
	apiHandler := api.NewAPI(BookService)
	
	apiHandler.Run(":8080")
	
}