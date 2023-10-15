package api

import (
	"net/http"
	"strconv"

	"github.com/brandoniu/lets-go/models"
	"github.com/brandoniu/lets-go/service"
	"github.com/gin-gonic/gin"
)

type API struct {
	Service service.BookService
	Router *gin.Engine
}

func NewAPI(s service.BookService) *API {
	router := gin.Default()
	api := &API{
		Service: s,
		Router: router,
	}
	api.setupRoutes()
	return api
}

func (api *API) setupRoutes() {
	api.Router.GET("/books", api.GetAllBooksHandler)
	api.Router.GET("/books/:id", api.GetBookByIDHandler)
	api.Router.POST("/books", api.AddBookHandler)
	api.Router.PUT("/books/:id", api.UpdateBookHandler)   
	api.Router.DELETE("/books/:id", api.DeleteBookHandler)
}

func (api *API) Run(port string) {
	api.Router.Run(port)
}

func (api *API) GetRouter() *gin.Engine {
	return api.Router
}

func (api *API) GetAllBooksHandler(c *gin.Context) {
	books, err := api.Service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (api *API) GetBookByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}
	book, err := api.Service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving the book"})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}
	c.JSON(http.StatusOK, book)
}

func (api *API) AddBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	id, err := api.Service.AddBook(&book)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding the book"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (api *API) UpdateBookHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
			return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	book.ID = int(id)

	err = api.Service.UpdateBook(&book)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the book"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book updated successfully"})
}

func (api *API) DeleteBookHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
			return
	}

	err = api.Service.RemoveBook(id)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting the book"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book deleted successfully"})
}



