package middlewares

import (
	"net/http"
	UserRepository "resedist/internal/modules/user/repositories"
	"resedist/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = UserRepository.New()
	return func(c *gin.Context) {

		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID != 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.Next()
	}
}
