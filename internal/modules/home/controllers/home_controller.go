package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	articleRepository "resedist/internal/modules/article/repositories"
)

type Controller struct {
	articleRepository articleRepository.ArticleRepositoryInterface
}

func New() *Controller {

	return &Controller{
		articleRepository: articleRepository.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	//html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	//	"title": "Home Page",
	//})
	c.JSON(http.StatusOK, gin.H{
		"articles": controller.articleRepository.List(8),
	})
}
