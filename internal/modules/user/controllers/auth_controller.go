package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resedist/internal/modules/user/requests/auth"
	"resedist/pkg/html"
)

type Controller struct{}

func New() *Controller {

	return &Controller{}
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

	// check errors
	// redirect

	c.JSON(http.StatusOK, gin.H{
		"message": "Register Done",
	})
}
