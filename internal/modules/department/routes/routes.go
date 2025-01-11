package routes

import (
	"resedist/internal/middlewares"
	depCtrl "resedist/internal/modules/department/controllers"
	depTypeCtrl "resedist/internal/modules/department/department_type/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	DepartmentController := depCtrl.New()
	DepartmentTypeController := depTypeCtrl.New()

	authGroup := router.Group("/api/v1")
	authGroup.GET("/department-type", DepartmentTypeController.Search)
	authGroup.Use(middlewares.IsAuthJwt())
	{
		authGroup.POST("/department", DepartmentController.Store)

	}

}
