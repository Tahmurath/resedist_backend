package routes

import (
	"github.com/gin-gonic/gin"
	userCtrl "resedist/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {

	UserController := userCtrl.New()
	router.GET("/register", UserController.Register)
	router.POST("/register", UserController.HandleRegister)

}
