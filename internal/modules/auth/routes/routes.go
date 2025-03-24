package routes

import (
	"resedist/internal/middlewares"
	authCtrl "resedist/internal/modules/auth/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	AuthController := authCtrl.New()

	router.GET("/auth/login", AuthController.Login)

	guestGroup := router.Group("/api/v1/auth")
	guestGroup.Use(middlewares.IsGuestJwt())
	{
		guestGroup.POST("/register", AuthController.HandleRegister)
		guestGroup.POST("/login", AuthController.HandleLogin)
	}

	authGroup := router.Group("/api/v1/auth")
	authGroup.Use(middlewares.IsAuthJwt())
	{
		authGroup.GET("/user", AuthController.User)
		authGroup.POST("/refresh", AuthController.RefreshAccessToken)
	}

}
