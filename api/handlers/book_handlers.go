package handlers

import (
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