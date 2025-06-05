package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resedist/pkg/html"
)

type Controller struct {
}

func New(router *gin.Engine) *Controller {
	return &Controller{}
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
