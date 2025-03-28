package routes

import (
	authRoutes "resedist/internal/modules/auth/routes"
	depRoutes "resedist/internal/modules/department/department/routes"
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

	userRoutes.Routes(router)
	authRoutes.Routes(router)
	depRoutes.Routes(router)

}

func RegisterSwaggerRoute(router *gin.Engine) {
	// Redirect از ریشه به Swagger
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func ConfigureCorsConfig(router *gin.Engine) {

	//gin.SetMode(gin.ReleaseMode)
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
