package controllers

import (
	"fmt"
	"net/http"
	configStruct "resedist/config"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/html"
	"resedist/pkg/rest"

	"github.com/gin-gonic/gin"
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
		"data":   string(data),
	})
}

func (ctl *Controller) ProtectedTG(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
