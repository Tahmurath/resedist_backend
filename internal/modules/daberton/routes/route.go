package routes

import (
	"github.com/gin-gonic/gin"

	"resedist/internal/middlewares"
	roomCtrl "resedist/internal/modules/daberton/controllers"
)

func Routes(router *gin.Engine) {
	roomController := roomCtrl.New()

	roomGroup := router.Group("/api/v1/daberton")

	//roomGroup.POST("/roomtemplate", roomController.CreateRoomTemplate)

	roomGroup.Use(middlewares.IsAuthJwt())
	{
		roomGroup.POST("/roomtemplate", roomController.CreateRoomTemplate)

		// authGroup.GET("/department-type", DepartmentTypeController.Search)

	}

	// tgGroup = router.Group("/api/tg/miniapp")
	// tgGroup.Use(middlewares.TgAuthMiddleware())
	// {
	// 	tgGroup.POST("/auth", tgController.TgAuth)
	// }

	//tgGroup.Use(middlewares.TgAuthMiddleware())
	//{
	//	tgGroup.POST("/protected", tgController.ProtectedTG)
	//}

}
