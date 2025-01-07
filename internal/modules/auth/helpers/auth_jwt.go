package helpers

import (
	"encoding/json"
	"net/http"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT(c *gin.Context) UserResponse.User {
	if user, exist := c.Get("user"); exist {
		if typedUser, ok := user.(UserResponse.User); ok {
			return typedUser
		}
	}

	cfg := config.Get()

	var response UserResponse.User

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return response
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		return response
	}

	tokenString := authHeaderParts[1]

	// Parse the token without claims validation
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Jwt.Secret), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return response
	}

	claims := token.Claims.(jwt.MapClaims)
	user := claims["sub"]

	userModelBytes, _ := json.Marshal(user)
	_ = json.Unmarshal(userModelBytes, &response)

	//Store the user in the context
	c.Set("user", response)

	return response
}
