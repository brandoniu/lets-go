package api

import (
	"github.com/brandoniu/lets-go/store"
	"github.com/gin-gonic/gin"
)

type API struct {
	Engine *gin.Engine
	Store *store.Store
}

func NewAPI(dbStore *store.Store) *API {
	api := &API{
		Engine: gin.Default(),
		Store: dbStore,
	}
	api.routes()

	return api
}

func (api *API) routes() {
	api.Engine.GET("/", api.health)
	api.Engine.GET("/health", api.health)
	api.Engine.GET("/books", api.getAllBooks)
}

func (api *API) getAllBooks(c *gin.Context) {
	books, err := api.Store.GetAllBooks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, books)
}

func (api *API) health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "Up" })
}