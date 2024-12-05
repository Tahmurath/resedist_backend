package routing

import (
	// "github.com/gin-contrib/cors"

	"github.com/gin-contrib/cors"
	"resedist/internal/providers/routes"
	"resedist/pkg/config"

	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"
	//"github.com/gin-contrib/cors"
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

func ConfigureCorsConfig() {
	router = gin.Default()
	cfg := config.Get()
	corsConfig := cors.Config{
		AllowOrigins:     cfg.Cors.AllowOrigins,
		AllowMethods:     cfg.Cors.AllowMethods,
		AllowHeaders:     cfg.Cors.AllowHeaders,
		ExposeHeaders:    cfg.Cors.ExposeHeaders,
		AllowCredentials: cfg.Cors.AllowCredentials,
	}
	router.Use(cors.New(corsConfig))
}
