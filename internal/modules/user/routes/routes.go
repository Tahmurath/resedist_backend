package routes

import (
	"resedist/internal/middlewares"
	userCtrl "resedist/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	UserController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/register", UserController.Register)
		guestGroup.POST("/register", UserController.HandleRegister)

		guestGroup.GET("/login", UserController.Login)
		guestGroup.POST("/login", UserController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", UserController.HandleLogout)
	}

}
