package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	configStruct "resedist/config"
	"resedist/internal/modules/user/requests/auth"
	UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/applog"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/html"
	"resedist/pkg/rest"
	"time"
)

var jwtKey = []byte("fc2e19d78c179b5dbb5358069f73156f835030ee43afe0fa9e257cdb421ccc5c")

type Claims struct {
	ID   uint
	Type string
	jwt.RegisteredClaims
}
type Controller struct {
	userService UserService.UserServiceInterface
	errFmt      *errors.ErrorFormat
	json        *rest.Jsonresponse
	rest        configStruct.Rest
}

func New(router *gin.Engine) *Controller {
	return &Controller{
		userService: UserService.New(),
		rest:        config.Get().Rest,
		errFmt:      errors.New(),
		json:        rest.New(),
	}
}

// @Summary Login and get JWT token
// @Description Authenticate user and return a JWT token
// @Tags SSO
// @Accept json
// @Produce json
// @Param user query auth.LoginRequest true "User data"
// @Success 200 {object} map[string]string "Token"
// @Router /sso/v1/auth/login [post]
func (ctl *Controller) HandleLogin(c *gin.Context) {
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
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

	access_token, err := GenerateAccessToken(user.ID)
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

	applog.Info("The user logged in successfully")

	c.SetCookie(
		"refresh_token",
		refresh_token,
		int(time.Now().Add(config.Get().Jwt.RefreshDuration).Unix()),
		"/",
		"localhost",
		true,
		true,
	)

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
// @Tags SSO
// @Accept json
// @Produce json
// @Param user query auth.RefreshRequest true "User data"
// @Success 200 {object} map[string]string "Token"
// @Router /sso/v1/auth/refresh [post]
func (ctl *Controller) RefreshAccessToken(c *gin.Context) {
	//var request auth.RefreshRequest
	//
	//if err := c.ShouldBind(&request); err != nil {
	//	ctl.json.Badrequest(c, rest.RestConfig{
	//		Error_message: ctl.errFmt.SetFromError(err),
	//	})
	//	return
	//}

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Http:          http.StatusUnauthorized,
			Error_message: err.Error(),
		})
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid || claims.Type != "refresh" {
		ctl.json.NotFound(c, rest.RestConfig{
			Http:          http.StatusUnauthorized,
			Error_message: err.Error(),
		})
		return
	}

	access_token, err := GenerateAccessToken(claims.ID)

	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{
			Error_message: err.Error(),
			Http:          http.StatusInternalServerError,
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: map[string]interface{}{
			"access_token": access_token,
		},
	})
}

func GenerateAccessToken(user_id uint) (string, error) {

	claims := &Claims{
		ID:   user_id,
		Type: "access",
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
		ID:   user.ID,
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Get().Jwt.RefreshDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func (ctl *Controller) Home(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/sso/html/home", gin.H{
		"title": "Create article",
	})
}

func (ctl *Controller) About(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/sso/html/about", gin.H{
		"title": "Create article",
	})
}
