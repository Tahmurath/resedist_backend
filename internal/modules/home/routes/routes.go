package routes

import (
	"github.com/gin-gonic/gin"
	homeCtrl "resedist/internal/modules/home/controllers"
)

func Routes(router *gin.Engine) {

	HomeController := homeCtrl.New()
	router.GET("/", HomeController.Index)

}
