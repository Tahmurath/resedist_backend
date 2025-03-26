package controllers

import (
	_ "net/http"
	"strings"
	"time"

	// "resedist/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// var jwtKey = []byte(config.Get().Jwt.Secret)
var jwtKey = []byte("fc2e19d78c179b5dbb5358069f73156f835030ee43afe0fa9e257cdb421ccc5c")

func generateAccessToken(userID uint) (string, error) {

	// duration := config.Get().Jwt.AccessDuration
	duration := 15 * time.Minute
	claims := &Claims{
		UserID: userID,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func generateRefreshToken(userID uint) (string, error) {

	// duration := config.Get().Jwt.RefreshDuration
	duration := 7 * 24 * time.Hour

	claims := &Claims{
		UserID: userID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func authMiddleware(tokenType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "توکن لازم است"})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "فرمت توکن نامعتبر"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid || claims.Type != tokenType {
			c.JSON(401, gin.H{"error": "توکن نامعتبر"})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func loginHandler(c *gin.Context) {
	var loginData struct {
		UserID      uint `json:"user_id" binding:"required"`
		NeedRefresh bool `json:"need_refresh"` // آیا Refresh Token هم بده؟
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "داده نامعتبر"})
		return
	}

	accessToken, err := generateAccessToken(loginData.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "خطا در تولید توکن"})
		return
	}

	response := gin.H{"access_token": accessToken}
	if loginData.NeedRefresh {
		refreshToken, err := generateRefreshToken(loginData.UserID)
		if err != nil {
			c.JSON(500, gin.H{"error": "خطا در تولید توکن رفرش"})
			return
		}
		response["refresh_token"] = refreshToken
	}
	c.JSON(200, response)
}

func refreshHandler(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "داده نامعتبر"})
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(input.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid || claims.Type != "refresh" {
		c.JSON(401, gin.H{"error": "توکن رفرش نامعتبر"})
		return
	}

	accessToken, err := generateAccessToken(claims.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "خطا در تولید توکن"})
		return
	}
	c.JSON(200, gin.H{"access_token": accessToken})
}

func getDepartmentHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")
	c.JSON(200, gin.H{"message": "Department retrieved", "user_id": userID})
}

func main() {
	r := gin.Default()

	r.POST("/login", loginHandler)
	r.POST("/refresh", refreshHandler)

	protected := r.Group("/api")
	protected.Use(authMiddleware("access"))
	{
		protected.GET("/department", getDepartmentHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
