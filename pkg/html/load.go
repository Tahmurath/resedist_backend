package html

import "github.com/gin-gonic/gin"

func LoadHTML(router *gin.Engine) {
	// internals/modules/moduleName/html/view.tmpl
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
