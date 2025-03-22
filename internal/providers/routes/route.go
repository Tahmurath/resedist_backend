package routes

import (
	articleRoutes "resedist/internal/modules/article/routes"
	authRoutes "resedist/internal/modules/auth/routes"
	depRoutes "resedist/internal/modules/department/department/routes"
	homeRoutes "resedist/internal/modules/home/routes"
	userRoutes "resedist/internal/modules/user/routes"

	"github.com/gin-gonic/gin"

	_ "resedist/docs"

	swaggerFiles "github.com/swaggo/files"
	//swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	ginSwagger "github.com/swaggo/gin-swagger"

	"resedist/pkg/config"

	"github.com/gin-contrib/cors"
)

func RegisterRoutes(router *gin.Engine) {

	ConfRoutes(router)

	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
	authRoutes.Routes(router)
	depRoutes.Routes(router)

}

func ConfRoutes(router *gin.Engine) {

	cfg := config.Get()
	corsConfig := cors.Config{
		AllowOrigins:     cfg.Cors.AllowOrigins,
		AllowMethods:     cfg.Cors.AllowMethods,
		AllowHeaders:     cfg.Cors.AllowHeaders,
		ExposeHeaders:    cfg.Cors.ExposeHeaders,
		AllowCredentials: cfg.Cors.AllowCredentials,
	}
	router.Use(cors.New(corsConfig))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
