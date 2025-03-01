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

	authGroup.GET("/department/:id", DepartmentController.Show)

	authGroup.Use(middlewares.IsAuthJwt())
	{
		authGroup.GET("/department-type2", DepartmentTypeController.Search)

		authGroup.GET("/department-type", DepartmentTypeController.Search)
		authGroup.GET("/department", DepartmentController.Search)
		authGroup.POST("/department", DepartmentController.Store)

	}

}
