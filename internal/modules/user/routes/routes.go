package routes

import (
	userCtrl "resedist/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	UserController := userCtrl.New()
	router.GET("/register", UserController.Register)
	router.POST("/register", UserController.HandleRegister)

	router.GET("/login", UserController.Login)
	router.POST("/login", UserController.HandleLogin)

}
