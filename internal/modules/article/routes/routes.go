package routes

import (
	"resedist/internal/middlewares"
	articleCtrl "resedist/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	ArticleController := articleCtrl.New()

	authGroup := router.Group("/articles")

	authGroup.GET("/:id", ArticleController.Show)

	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", ArticleController.Create)
	}

}
