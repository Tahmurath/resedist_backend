package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Register Done",
	})
}
