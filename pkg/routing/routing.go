package routing

import (
	// "github.com/gin-contrib/cors"

	"resedist/internal/providers/routes"
	// "resedist/pkg/config"

	// "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"
	//"github.com/gin-contrib/cors"
	// _ "resedist/docs"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	"resedist/pkg/config"
)

func Init() {
	SetGinMode()
	router = gin.Default()
	SetTrustedProxies(router)
}

func GetRouter() *gin.Engine {
	return router
}

func SetTrustedProxies(router *gin.Engine) {
	router.SetTrustedProxies(config.Get().Server.TrustedProxies)
}

func SetGinMode() {
	// configs :=
	if gin.Mode() != gin.ReleaseMode {
		gin.SetMode(config.Get().Server.Ginmode)
	}
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}

func ConfigureCorsConfig() {
	routes.ConfigureCorsConfig(GetRouter())
}

func RegisterSwaggerRoute() {
	routes.RegisterSwaggerRoute(GetRouter())
}
