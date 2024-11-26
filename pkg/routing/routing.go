package routing

import (
	// "github.com/gin-contrib/cors"

	"resedist/internal/providers/routes"

	"github.com/gin-gonic/gin"

	// cors "github.com/rs/cors/wrapper/gin"
	"github.com/gin-contrib/cors"
)

func Init() {
	router = gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // یا "*" برای اجازه به همه دامنه‌ها
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))
	// router.Use(cors.Default())
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
