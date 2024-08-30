package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resedist/pkg/html"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {

		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title": "Home Page",
		})
	})

	router.GET("/about", func(c *gin.Context) {

		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "about Page",
		})

	})
}
