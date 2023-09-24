package api

import (
	"github.com/brandoniu/lets-go/api/handlers"
	"github.com/brandoniu/lets-go/api/routes"
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
	h := handlers.NewHandlers(api.Store)
	routes.SetupRoutes(api.Engine, h)

	return api
}
