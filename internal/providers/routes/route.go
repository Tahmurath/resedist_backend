package routes

import (
	articleRoutes "resedist/internal/modules/article/routes"
	homeRoutes "resedist/internal/modules/home/routes"
	userRoutes "resedist/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)

}
