package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	configStruct "resedist/config"
	tgUserServices "resedist/internal/modules/tgminiapp/services"
	"resedist/internal/modules/user/requests/auth"
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

	tg_id := c.GetInt64("tg_user_id")
	tg_username := c.GetString("tg_user_name")

	//fmt.Println(tg_id.(int64))
	//return

	if ctl.tgUserService.CheckUserExist(tg_id) {

		ctl.json.Success(c, rest.RestConfig{
			Data: map[string]interface{}{
				"message": "User exist",
				"tg_id":   tg_id,
			},
		})

	} else {

		create, err := ctl.UserService.Create(auth.RegisterRequest{
			Name: fmt.Sprint(tg_username),
			//Email:    fmt.Sprint(tg_username) + "@" + fmt.Sprint(tg_id) + ".com",
			//Password: "tguser" + fmt.Sprint(tg_id),
		})

		if err != nil {
			ctl.json.ServerError(c, rest.RestConfig{
				Error_message: err.Error(),
				Http:          http.StatusInternalServerError,
			})
			return
		}

		//tgcreate, err := ctl.tgUserService.Create(tg_id, create.ID)
		//if err != nil {
		//	ctl.json.ServerError(c, rest.RestConfig{
		//		Error_message: err.Error(),
		//		Http:          http.StatusInternalServerError,
		//	})
		//	return
		//}

		ctl.json.Success(c, rest.RestConfig{
			Data: map[string]interface{}{
				"message": "User created",
				"tg_id":   tg_id,
				"user":    create,
			},
		})
		//ctl.UserService.Create(UserResponse.User{
		//	Name:  fmt.Sprint(tg_username),
		//	Email: fmt.Sprint(tg_username) + "@" + fmt.Sprint(tg_id) + ".com",
		//	Password: "tguser" + fmt.Sprint(tg_id),
		//	//TgId:     tg_id,
		//})

		//ctl.tgUserService.

	}
	//if user, exist := c.Get("user"); exist {
	//	if typedUser, ok := user.(UserResponse.User); ok {
	//		return typedUser
	//	}
	//}
	//fmt.Println(tg_user)

	//ctl.json.Success(c, rest.RestConfig{
	//	Data: map[string]interface{}{
	//		"user": tg_id,
	//	},
	//})
}
