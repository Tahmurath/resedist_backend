package helpers

import (
	"github.com/gin-gonic/gin"
	UserRepository "resedist/internal/modules/user/repositories"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/sessions"
	"strconv"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User
	authID := sessions.Get(c, "auth")

	userID, _ := strconv.Atoi(authID)

	var userRepo = UserRepository.New()
	user := userRepo.FindByID(userID)

	if user.ID == 0 {
		return response
	}

	return UserResponse.ToUser(user)

}
