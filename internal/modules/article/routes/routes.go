package routes

import (
	"github.com/gin-gonic/gin"
	articleCtrl "resedist/internal/modules/article/controllers"
)

func Routes(router *gin.Engine) {

	ArticleController := articleCtrl.New()
	router.GET("/articles/:id", ArticleController.Show)

}
