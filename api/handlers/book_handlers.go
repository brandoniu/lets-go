package handlers

import (
	"fmt"
	"strconv"

	"github.com/brandoniu/lets-go/store"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Store *store.Store
}

func NewHandlers(s *store.Store) *Handlers {
	return &Handlers{Store: s}
}

func (h *Handlers) GetAllBooks(c *gin.Context) {
	books, err := h.Store.GetAllBooks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, books)
}

func (h *Handlers) Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "Up" })
}

// ...

func (h *Handlers) CreateBook(c *gin.Context) {
	var book store.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	if err := h.Store.CreateBook(&book); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, book)
}

func (h *Handlers) UpdateBook(c *gin.Context) {
	var book store.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	if err := h.Store.UpdateBook(&book); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, book)
}

func (h *Handlers) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	s, err := strconv.Atoi(id);
	if err != nil {
		fmt.Println("Invalid number")
	}
	if err := h.Store.DeleteBook(s); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Book deleted successfully"})
}
