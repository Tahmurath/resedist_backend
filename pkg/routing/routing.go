package routing

import (
	// "github.com/gin-contrib/cors"

	"resedist/internal/providers/routes"
	// "resedist/pkg/config"

	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"
	//"github.com/gin-contrib/cors"
	// _ "resedist/docs"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
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

// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization

func ConfigureCorsConfig() {
	router = gin.Default()
	// cfg := config.Get()
	// corsConfig := cors.Config{
	// 	AllowOrigins:     cfg.Cors.AllowOrigins,
	// 	AllowMethods:     cfg.Cors.AllowMethods,
	// 	AllowHeaders:     cfg.Cors.AllowHeaders,
	// 	ExposeHeaders:    cfg.Cors.ExposeHeaders,
	// 	AllowCredentials: cfg.Cors.AllowCredentials,
	// }
	// router.Use(cors.New(corsConfig))
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
