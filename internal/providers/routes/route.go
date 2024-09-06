package routes

import (
	"github.com/gin-gonic/gin"
	articleRoutes "resedist/internal/modules/article/routes"
	homeRoutes "resedist/internal/modules/home/routes"
	userRoutes "resedist/internal/modules/user/routes"
)

func RegisterRoutes(router *gin.Engine) {

	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)

}
