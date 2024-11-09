package routes

import (
	homeCtrl "resedist/internal/modules/home/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	HomeController := homeCtrl.New()
	router.GET("/", HomeController.Index)

}
