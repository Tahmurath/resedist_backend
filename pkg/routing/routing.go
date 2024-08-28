package routing

import (
	"github.com/gin-gonic/gin"
	"resedist/internal/providers/routes"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
