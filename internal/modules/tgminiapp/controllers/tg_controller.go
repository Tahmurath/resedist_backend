package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/http"
	configStruct "resedist/config"
	tgAuth "resedist/internal/modules/tgminiapp/requests/auth"
	TgUserResponse "resedist/internal/modules/tgminiapp/responses"

	//tgUserResponse "resedist/internal/modules/tgminiapp/responses"
	tgUserServices "resedist/internal/modules/tgminiapp/services"
	"resedist/internal/modules/user/requests/auth"
	//UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/html"
	"resedist/pkg/jwtutil"
	"resedist/pkg/rest"
)

var jwtKey = []byte("fc2e19d78c179b5dbb5358069f73156f835030ee43afe0fa9e257cdb421ccc5c")

type Controller struct {
	tgUserService tgUserServices.TgUserServiceInterface
	UserService   UserService.UserServiceInterface
	errFmt        *errors.ErrorFormat
	json          *rest.Jsonresponse
	rest          configStruct.Rest
}

func New() *Controller {

	return &Controller{
		//departmentService: DepartmentService.New(),
		tgUserService: tgUserServices.New(),
		UserService:   UserService.New(),
		rest:          config.Get().Rest,
		errFmt:        errors.New(),
		json:          rest.New(),
	}
}

// @Summary Tg miniapp auth
// @Description Returns a JWT token for authenticated user
// @Security BearerAuth
// @ID tg-miniapp-auth
// @Tags tgminiapp
// @Accept json
// @Produce json
// @Param tg_miniapp_auth query string true "Tg miniapp auth data"
// @Success 200 {object} map[string]string "Token"
// @Router /api/v1/tgminiapp/auth [get]
func (ctl *Controller) TelegramMiniAppIndex(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/tgminiapp/html/miniapp", gin.H{
		"title": "Create article",
	})
	// bottoken := config.Get().Telegram.BotToken
	// c.JSON(200, gin.H{
	// 	"message": "tg miniapp auth" + bottoken,
	// })
}

func (ctl *Controller) TelegramCallBack(c *gin.Context) {

	data, _ := c.GetRawData()
	fmt.Println("Telegram Callback:", string(data))
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		//"data":   string(data),
	})
}

func (ctl *Controller) TgAuth(c *gin.Context) {

	tg_user, _ := c.Get("tg_user")
	tgUser, ok := tg_user.(initdata.InitData)
	if !ok {
		ctl.json.ServerError(c, rest.RestConfig{Error_message: "Invalid user data"})
		return
	}

	user, err := ctl.getOrCreateUser(tgUser)
	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{Error_message: err.Error()})
		return
	}

	accessToken, refreshToken, err := ctl.generateTokens(user.ID)
	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{Error_message: err.Error()})
		return
	}

	ctl.setTokenCookie(c, refreshToken, "refresh_token")
	ctl.setTokenCookie(c, accessToken, "access_token")

	ctl.json.Success(c, rest.RestConfig{
		Data: map[string]interface{}{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"user":          user,
		},
	})
}

func (ctl *Controller) generateTokens(userID uint) (string, string, error) {
	accessToken, err := jwtutil.GenerateAccessToken(userID, "tgminiapp")
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwtutil.GenerateRefreshToken(userID, "tgminiapp")
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (ctl *Controller) setTokenCookie(c *gin.Context, jwtToken string, name string) {
	c.SetCookie(
		name,
		jwtToken,
		int(config.Get().Jwt.RefreshDuration.Seconds()),
		"/",
		"abler-carmela-pliant.ngrok-free.dev",
		true,
		true,
	)
}

func (ctl *Controller) getOrCreateUser(tgUser initdata.InitData) (TgUserResponse.TgUser, error) {

	var telegramUser TgUserResponse.TgUser

	if user, found := ctl.tgUserService.FindByTgID(tgUser.User.ID); found {
		return user, nil
	}

	user, err := ctl.UserService.Create(auth.RegisterRequest{Name: fmt.Sprint(tgUser.User.Username)})
	if err != nil {
		return telegramUser, err
	}

	telegramUser, err = ctl.tgUserService.Create(tgAuth.TgRegisterRequest{
		TgID:         tgUser.User.ID,
		FirstName:    tgUser.User.FirstName,
		LastName:     tgUser.User.LastName,
		Username:     tgUser.User.Username,
		LanguageCode: tgUser.User.LanguageCode,
		PhotoURL:     tgUser.User.PhotoURL,
		IsBot:        tgUser.User.IsBot,
		IsPremium:    tgUser.User.IsPremium,
	}, user)

	return telegramUser, err
}

func (ctl *Controller) RefreshAccessToken(c *gin.Context) {

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Http:          http.StatusUnauthorized,
			Error_message: err.Error(),
		})
		return
	}

	claims := &jwtutil.Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid || claims.Type != "refresh" || claims.ClientType != "tgminiapp" {
		ctl.json.NotFound(c, rest.RestConfig{
			Http:          http.StatusUnauthorized,
			Error_message: err.Error(),
		})
		return
	}

	accessToken, err := jwtutil.GenerateAccessToken(claims.ID, claims.ClientType)

	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{
			Error_message: err.Error(),
			Http:          http.StatusInternalServerError,
		})
		return
	}

	ctl.setTokenCookie(c, accessToken, "access_token")

	ctl.json.Success(c, rest.RestConfig{
		Data: map[string]interface{}{
			"access_token": accessToken,
		},
	})
}
