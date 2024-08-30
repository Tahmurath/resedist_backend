package html

import (
	"github.com/gin-gonic/gin"
	"resedist/internal/providers/view"
)

func Render(c *gin.Context, code int, name string, data gin.H) {

	data = view.WithGlobalData(data)
	c.HTML(code, name, data)
}
