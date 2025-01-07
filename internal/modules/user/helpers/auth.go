package helpers

import (
	UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) UserResponse.User {

	if user, exist := c.Get("user"); exist {
		if typedUser, ok := user.(UserResponse.User); ok {
			return typedUser
		}
	}

	var response UserResponse.User

	authID := sessions.Get(c, "auth")
	if authID == "" {
		return response
	}

	userID, err := strconv.Atoi(authID)
	if err != nil {
		return response
	}

	userService := UserService.New()
	foundUser, err := userService.GetCachedUserById(userID)
	if err != nil {
		return response
	}

	if foundUser.ID == 0 {
		return response
	}

	//convertedUser := UserResponse.ToUser(foundUser)
	c.Set("user", foundUser)

	return foundUser

}
