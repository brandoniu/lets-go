package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/brandoniu/lets-go/store"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable password=mypassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	store := store.New(db)

	r := gin.Default()

	r.GET("/books", func(c *gin.Context) {
		books, err := store.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
	})

	r.Run(":8080")

	
}