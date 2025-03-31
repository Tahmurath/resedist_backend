package routes

import (
	"resedist/internal/middlewares"
	depTypeCtrl "resedist/internal/modules/department/department_type/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	DepTypeController := depTypeCtrl.New()

	authGroup := router.Group("/api/v1")

	authGroup.Use(middlewares.IsAuthJwt())
	{
		authGroup.GET("/department-type", DepTypeController.Search)
		authGroup.GET("/department-type/:id", DepTypeController.Show)
		authGroup.POST("/department-type", DepTypeController.Store)
		authGroup.PUT("/department-type/:id", DepTypeController.Update)
		authGroup.DELETE("/department-type/:id", DepTypeController.Remove)

	}

}
