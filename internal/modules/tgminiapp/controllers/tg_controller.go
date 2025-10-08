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

func (ctl *Controller) TelegramMiniAppAuth(c *gin.Context) {
	user_id := c.GetString("user_id")
	username := c.GetString("user_id")
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"user_id":   user_id,
		"user_name": username,
	})
}
func (ctl *Controller) ProtectedTG(c *gin.Context) {
	//data, _ := c.GetRawData()
	//fmt.Println("Telegram Callback:", string(data))
	user_id, _ := c.Get("user_id")
	username, _ := c.Get("username")

	fmt.Println("Authenticated user_id:", user_id)
	fmt.Println("Authenticated username:", username)
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"user_id":   user_id,
		"user_name": username,
	})
}
