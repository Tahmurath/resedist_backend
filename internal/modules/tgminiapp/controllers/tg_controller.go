package controllers

import (
	"github.com/gin-gonic/gin"

	configStruct "resedist/config"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/rest"
)

type Controller struct {
	//departmentService DepartmentService.DepartmentServiceInterface
	errFmt *errors.ErrorFormat
	json   *rest.Jsonresponse
	rest   configStruct.Rest
}

func New() *Controller {

	return &Controller{
		//departmentService: DepartmentService.New(),
		rest:   config.Get().Rest,
		errFmt: errors.New(),
		json:   rest.New(),
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
func (ctl *Controller) TelegramMiniAppAuth(c *gin.Context) {

	bottoken := config.Get().Telegram.BotToken
	c.JSON(200, gin.H{
		"message": "tg miniapp auth" + bottoken,
	})
}
