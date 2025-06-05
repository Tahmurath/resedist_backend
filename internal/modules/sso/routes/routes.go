package routes

import (
	ssoCtrl "resedist/internal/modules/sso/controllers"

	// depTypeCtrl "resedist/internal/modules/department/department_type/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	ssoController := ssoCtrl.New(router)
	// DepartmentTypeController := depTypeCtrl.New()

	authGroup := router.Group("/sso")
	authGroup.GET("/about", ssoController.About)
	authGroup.GET("/home", ssoController.Home)

}
