package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resedist/internal/modules/user/requests/auth"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/html"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {

	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate
	var request auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}
	// create user
	user, err := controller.userService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// redirect
	log.Printf("user created with name %s", user.Name)
	c.Redirect(http.StatusFound, "/")

}
