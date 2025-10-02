package routing

import (
	// "github.com/gin-contrib/cors"

	"fmt"
	"log"
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

var router *gin.Engine

func Init() {
	SetGinMode()
	router = gin.Default()
	SetTrustedProxies(router)
}

func GetRouter() *gin.Engine {
	return router
}

func Serve() {
	r := GetRouter()

	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// log.Fatal(autotls.Run(r, "localhost"))
	//log.Fatal(http.ListenAndServeTLS("localhost:4000", "localhost.crt", "localhost.key", r))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}

//openssl req -new -subj "/C=US/ST=Utah/CN=localhost" -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr
// openssl x509 -req -days 365 -in localhost.csr -signkey localhost.key -out localhost.crt
//log.Fatal(http.ListenAndServeTLS("localhost:4000", "localhost.crt", "localhost.key", r))

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
