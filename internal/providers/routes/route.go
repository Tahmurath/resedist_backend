package routes

import (
	"github.com/gin-gonic/gin"
	homeRoutes "resedist/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}
