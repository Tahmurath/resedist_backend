package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	configStruct "resedist/config"
	//tgUserResponse "resedist/internal/modules/tgminiapp/responses"
	tgUserServices "resedist/internal/modules/tgminiapp/services"
	//UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/html"
	"resedist/pkg/rest"
)

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

	html.Render(c, http.StatusOK, "modules/tgminiapp/html/miniapp2", gin.H{
		"title": "Create article",
	})
	// bottoken := config.Get().Telegram.BotToken
	// c.JSON(200, gin.H{
	// 	"message": "tg miniapp auth" + bottoken,
	// })
}
