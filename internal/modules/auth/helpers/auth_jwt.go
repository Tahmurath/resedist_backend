package helpers

import (
	"encoding/json"
	"net/http"
	UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/applog"
	"resedist/pkg/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getTokenFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return ""
	}

	return authHeaderParts[1]
}

func getTokenFromCookie(c *gin.Context) string {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return ""
	}
	return cookie
}

func getTokenFromQuery(c *gin.Context) string {
	return c.Query("refresh_token")
}

func getUserByToken(tokenString string) (UserResponse.User, error) {
	var response UserResponse.User
	cfg := config.Get()

	// Parse the token without claims validation
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Jwt.Secret), nil
	})

	if err != nil || !token.Valid {
		return response, err
	}

	claims := token.Claims.(jwt.MapClaims)
	user := claims

	userModelBytes, _ := json.Marshal(user)
	_ = json.Unmarshal(userModelBytes, &response)

	return response, nil
}

func getUserFromContext(c *gin.Context) (UserResponse.User, bool) {
	if user, exist := c.Get("user"); exist {
		if typedUser, ok := user.(UserResponse.User); ok {
			return typedUser, true
		}
	}
	return UserResponse.User{}, false
}

func AuthJWT(c *gin.Context) UserResponse.User {
	// Try to get user from context

	var response UserResponse.User

	if user, found := getUserFromContext(c); found {
		applog.Info("return user from context")
		return user
	}

	// Try to get token from header
	tokenString := getTokenFromHeader(c)
	if tokenString == "" {
		// Try to get token from cookie
		tokenString = getTokenFromCookie(c)
	}
	if tokenString == "" {
		// Try to get token from query parameter
		tokenString = getTokenFromQuery(c)
	}
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
		return response
	}

	user, err := getUserByToken(tokenString)
	if err != nil || user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return response
	}

	userService := UserService.New()
	foundUser, err := userService.GetCachedUserById(int(user.ID))
	if err != nil {
		return response
	}

	if foundUser.ID == 0 {
		return response
	}

	// Store the user in the context
	c.Set("user", user)
	applog.Info("return user by parse token and set into context")

	return user
}

//func AuthJWT(c *gin.Context) UserResponse.User {
//	if user, exist := c.Get("user"); exist {
//		if typedUser, ok := user.(UserResponse.User); ok {
//			applog.Info("return user from contxt")
//			return typedUser
//		}
//	}
//
//	cfg := config.Get()
//
//	var response UserResponse.User
//
//	authHeader := c.GetHeader("Authorization")
//	if authHeader == "" {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
//		return response
//	}
//
//	authHeaderParts := strings.Split(authHeader, " ")
//	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
//		return response
//	}
//
//	tokenString := authHeaderParts[1]
//
//	// Parse the token without claims validation
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(cfg.Jwt.Secret), nil
//	})
//
//	if err != nil || !token.Valid {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//		return response
//	}
//
//	claims := token.Claims.(jwt.MapClaims)
//	user := claims
//
//	userModelBytes, _ := json.Marshal(user)
//	_ = json.Unmarshal(userModelBytes, &response)
//
//	//fmt.Println("auth_jwt->response:", response)
//
//	//Store the user in the context
//	c.Set("user", response)
//	applog.Info("return user by parse jwt and set into context")
//
//	return response
//}
