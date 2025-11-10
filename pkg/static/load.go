package static

import "github.com/gin-gonic/gin"

func LoadStatic(router *gin.Engine) {
	router.Static("/assets", "./assets/assets")
	router.Static("/js", "./assets/js")
	router.Static("/locales", "./assets/locales")
	router.Static("/front", "./front")
}
