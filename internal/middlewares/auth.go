package middlewares

import (
	"net/http"
	authHelpers "resedist/internal/modules/auth/helpers"
	userHelpers "resedist/internal/modules/user/helpers"

	"github.com/gin-gonic/gin"
)

func IsAuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := authHelpers.AuthJWT(c)

		if user.ID == 0 {

			return
		}

		c.Next()
	}
}

func IsAuth() gin.HandlerFunc {
	//var userRepo = UserRepository.New()

	return func(c *gin.Context) {

		user := userHelpers.Auth(c)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Next()
	}
}
