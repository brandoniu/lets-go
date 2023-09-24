package routes

import (
	"github.com/brandoniu/lets-go/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handlers) {
	setupBookRoutes(r, h)
}