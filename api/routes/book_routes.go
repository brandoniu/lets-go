package routes

import (
	"github.com/brandoniu/lets-go/api/handlers"
	"github.com/gin-gonic/gin"
)

func setupBookRoutes(r *gin.Engine, h *handlers.Handlers) {
	r.GET("/", h.Health)
	r.GET("/health", h.Health)
	r.GET("/books", h.GetAllBooks)
	r.POST("/books", h.CreateBook)
	r.PUT("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)

}
