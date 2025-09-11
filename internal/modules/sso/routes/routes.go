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

	staticGroup := router.Group("/sso")
	staticGroup.GET("/about", ssoController.About)
	staticGroup.GET("/home", ssoController.Home)

	//staticGroup.POST("/v1/auth/refresh", ssoController.RefreshAccessToken)

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

}
