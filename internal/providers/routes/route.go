package routes

import (
	articleRoutes "resedist/internal/modules/article/routes"
	authRoutes "resedist/internal/modules/auth/routes"
	depRoutes "resedist/internal/modules/department/routes"
	homeRoutes "resedist/internal/modules/home/routes"
	userRoutes "resedist/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
	authRoutes.Routes(router)
	depRoutes.Routes(router)

}
