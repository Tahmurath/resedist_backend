package routes

import (
	"resedist/internal/middlewares"
	depCtrl "resedist/internal/modules/department/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	DepartmentController := depCtrl.New()

	authGroup := router.Group("/api/v1/department")
	authGroup.Use(middlewares.IsAuthJwt())
	{
		authGroup.POST("/new", DepartmentController.Store)
	}

}
