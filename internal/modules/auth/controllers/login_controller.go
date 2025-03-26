package controllers

import (
	"log"
	"net/http"
	"resedist/internal/modules/user/requests/auth"
	UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/applog"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/rest"

	configStruct "resedist/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	// UserService "resedist/internal/modules/user/services"
	"time"
)

var jwtKey = []byte("fc2e19d78c179b5dbb5358069f73156f835030ee43afe0fa9e257cdb421ccc5c")

type Claims struct {
	UserID uint   `json:"user_id"`
	Type   string `json:"type"` // "access" یا "refresh"
	Pack   interface{}
	jwt.RegisteredClaims
}
type Controller struct {
	userService UserService.UserServiceInterface
	errFmt      *errors.ErrorFormat
	json        *rest.Jsonresponse
	rest        configStruct.Rest
}

func New() *Controller {

	return &Controller{
		userService: UserService.New(),
		rest:        config.Get().Rest,
		errFmt:      errors.New(),
		json:        rest.New(),
	}
}

func (controller *Controller) Login(c *gin.Context) {
	//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	c.JSON(http.StatusOK, gin.H{
		"sdfsdfs": "sdfsdfs",
		"sdsda":   "sdfsdfs",
		"sdfsd":   "sdfsdfs",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromError(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with bind",
			"errors":  errors.Get(),
		})
		return
	}

	if controller.userService.CheckUserExist(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with email",
			"errors":  errors.Get(),
		})
		return
	}

	// Create the user
	user, err := controller.userService.Create(request)

	// Check if there is any error on the user creation
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with user creation",
		})
		return
	}

	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

// @Summary Login and get JWT token
// @Description Authenticate user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user query auth.LoginRequest true "User data"
// @Success 200 {object} map[string]string "Token"
// @Router /api/v1/auth/login [post]
func (ctl *Controller) HandleLogin(c *gin.Context) {
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			//Error_message: ctl.errFmt.SetFromError(err),
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user, err := ctl.userService.HandleUserLogin(request)
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	access_token, err := ctl.generateAccessToken(user.ID)
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
			Http:          http.StatusInternalServerError,
		})
		return
	}

	refresh_token, err := ctl.generateRefreshToken(user)
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
			Http:          http.StatusInternalServerError,
		})
		return
	}

	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie(
	// 	"refresh_token",
	// 	refresh_token,
	// 	int(time.Now().Add(config.Get().Jwt.RefreshDuration).Unix()),
	// 	"/",
	// 	"localhost:5173",
	// 	true,
	// 	true,
	// )

	applog.Info("The user logged in successfully")

	ctl.json.Success(c, rest.RestConfig{
		Data: map[string]interface{}{
			"access_token":  access_token,
			"refresh_token": refresh_token,
			"user":          user,
		},
	})
}

// @Summary Refresh access token
// @Description Refresh an access token using a refresh token
// @Security BearerAuth
// @Tags auth
// @Accept json
// @Produce json
// @Param user query auth.RefreshRequest true "User data"
// @Success 200 {object} map[string]string "Token"
// @Router /api/v1/auth/refresh [post]
func (ctl *Controller) RefreshAccessToken(c *gin.Context) {
	var request auth.RefreshRequest

	// refresh, err := c.Cookie("refresh_token")
	// if err != nil {
	// 	ctl.json.NotFound(c, rest.RestConfig{
	// 		Error_message: err.Error(),
	// 	})
	// 	return
	// }

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(request.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid || claims.Type != "refresh" {
		ctl.json.NotFound(c, rest.RestConfig{
			Http:          http.StatusUnauthorized,
			Error_message: err.Error(),
		})
		return
	}

	access_token, err := ctl.generateAccessToken(claims.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "خطا در تولید توکن"})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: map[string]interface{}{
			"access_token": access_token,
		},
	})

}

func (ctl *Controller) User(c *gin.Context) {
	user, _ := c.Get("auth")
	c.JSON(http.StatusOK, gin.H{"message": "Authenticated", "user": user})
}

func (ctl *Controller) generateAccessToken(user_id uint) (string, error) {

	claims := &Claims{
		UserID: user_id,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Get().Jwt.AccessDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Get().Jwt.Secret))
}

func (ctl *Controller) generateRefreshToken(user UserResponse.User) (string, error) {

	claims := &Claims{
		UserID: user.ID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Get().Jwt.RefreshDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
