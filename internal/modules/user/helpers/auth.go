package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	UserRepository "resedist/internal/modules/user/repositories"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/sessions"
	"strconv"
)

func Auth(c *gin.Context) UserResponse.User {

	if user, exist := c.Get("user"); exist {
		if typedUser, ok := user.(UserResponse.User); ok {
			fmt.Println(typedUser)
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

	userRepo := UserRepository.New()
	foundUser := userRepo.FindByID(userID)

	if foundUser.ID == 0 {
		return response
	}

	convertedUser := UserResponse.ToUser(foundUser)
	c.Set("user", convertedUser)

	return convertedUser

}
