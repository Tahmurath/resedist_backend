package middlewares

import (
	"encoding/json"
	"net/http"
	"resedist/internal/modules/user/helpers"
	"resedist/internal/modules/user/models"
	"resedist/internal/modules/user/responses"
	"resedist/pkg/config"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

func IsAuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()

		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		// Check if the header has the "Bearer" scheme
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			return
		}

		tokenString := authHeaderParts[1]

		// Parse the token without claims validation
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Jwt.Secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		user := claims["sub"]

		userModel := models.User{}
		userModelBytes, _ := json.Marshal(user)
		_ = json.Unmarshal(userModelBytes, &userModel)

		//Store the user in the context
		c.Set("auth", responses.ToUser(userModel))

		c.Next()
	}
}

func IsAuth() gin.HandlerFunc {
	//var userRepo = UserRepository.New()

	return func(c *gin.Context) {

		//authID := sessions.Get(c, "auth")
		//userID, _ := strconv.Atoi(authID)

		//user := userRepo.FindByID(userID)
		user := helpers.Auth(c)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		c.Next()
	}
}
