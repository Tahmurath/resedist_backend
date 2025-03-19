package routes

import (
	"resedist/internal/middlewares"
	depCtrl "resedist/internal/modules/department/department/controllers"
	depTypeCtrl "resedist/internal/modules/department/department_type/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	DepartmentController := depCtrl.New()
	DepartmentTypeController := depTypeCtrl.New()

	authGroup := router.Group("/api/v1")

	authGroup.GET("/department", DepartmentController.Search)

	authGroup.Use(middlewares.IsAuthJwt())
	{
		authGroup.GET("/department/:id", DepartmentController.Show)

		authGroup.POST("/department", DepartmentController.Store)
		authGroup.PUT("/department/:id", DepartmentController.Update)
		authGroup.DELETE("/department/:id", DepartmentController.Remove)

		authGroup.GET("/department-type", DepartmentTypeController.Search)

	}

}
