package routes

import (
	"github.com/gin-gonic/gin"
	"resedist/internal/middlewares"

	tgCtrl "resedist/internal/modules/tgminiapp/controllers"
)

func Routes(router *gin.Engine) {
	tgController := tgCtrl.New()

	tgGroup := router.Group("/tg/miniapp")

	tgGroup.GET("/", tgController.TelegramMiniAppIndex)
	tgGroup.POST("/callback", tgController.TelegramCallBack)

	tgGroup.Use(middlewares.TgAuthMiddleware())
	{
		tgGroup.POST("/auth", tgController.TgAuth)
	}

	//tgGroup.Use(middlewares.TgAuthMiddleware())
	//{
	//	tgGroup.POST("/protected", tgController.ProtectedTG)
	//}

}
