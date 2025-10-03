package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		//router.StaticFile()
		//html.Render(c, http.StatusOK, "modules/index/html/miniapp", gin.H{
		//	"title": "Create article",
		//})
		c.Redirect(302, "/swagger/index.html")
	})

	//func RegisterSwaggerRoute(router *gin.Engine) {
	//	// Redirect از ریشه به Swagger
	//	router.GET("/", func(c *gin.Context) {
	//		//router.StaticFile()
	//		html.Render(c, http.StatusOK, "modules/index/html/miniapp", gin.H{
	//			"title": "Create article",
	//		})
	//		//c.Redirect(302, "/swagger/index.html")
	//	})
	//	router.POST("/callback", func(c *gin.Context) {
	//		data, _ := c.GetRawData()
	//		fmt.Println("Telegram Callback:", string(data))
	//		c.JSON(http.StatusOK, gin.H{
	//			"status": "ok",
	//			"data":   string(data),
	//		})
	//	})
	//	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}
}
