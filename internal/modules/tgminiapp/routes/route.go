package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resedist/pkg/html"
)

func Routes(router *gin.Engine) {
	router.GET("/tgminiapp", func(c *gin.Context) {
		//router.StaticFile()
		html.Render(c, http.StatusOK, "modules/tgminiapp/html/miniapp", gin.H{
			"title": "Create article",
		})
		//c.Redirect(302, "/swagger/index.html")
	})
}
