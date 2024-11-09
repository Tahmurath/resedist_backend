package routing

import (
	// "github.com/gin-contrib/cors"

	"resedist/internal/providers/routes"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
)

func Init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
