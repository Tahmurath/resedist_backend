package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	//articleRepository "resedist/internal/modules/article/repositories"
	ArticleService "resedist/internal/modules/article/services"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {

	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	//html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	//	"title": "Home Page",
	//})
	c.JSON(http.StatusOK, gin.H{
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}
