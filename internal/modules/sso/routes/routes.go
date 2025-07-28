package routes

import (
	"resedist/internal/middlewares"
	ssoCtrl "resedist/internal/modules/sso/controllers"

	// depTypeCtrl "resedist/internal/modules/department/department_type/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	ssoController := ssoCtrl.New(router)
	// DepartmentTypeController := depTypeCtrl.New()

	//authGroup := router.Group("/sso")

	guestGroup := router.Group("/sso/v1/auth")
	guestGroup.Use(middlewares.IsGuestJwt())
	{
		//guestGroup.POST("/register", AuthController.HandleRegister)
		guestGroup.POST("/login", ssoController.HandleLogin)
	}

	authGroup := router.Group("/sso/v1/auth")
	authGroup.Use(middlewares.IsAuthJwt())
	{
		//authGroup.GET("/user", ssoController.User)
		authGroup.POST("/refresh", ssoController.RefreshAccessToken)
	}
	//authGroup.GET("/about", ssoController.About)
	//authGroup.GET("/home", ssoController.Home)

}
