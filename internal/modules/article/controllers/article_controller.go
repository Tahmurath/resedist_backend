package controllers

import (
	"net/http"
	"resedist/pkg/html"
	"strconv"

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

func (controller *Controller) Show(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{
			"title":   "Server Error",
			"message": "error converting the id",
		})
		return

	}
	article, err := controller.articleService.Find(id)

	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{
			"title":   "Page not found",
			"message": err.Error(),
		})
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{
		"title":   "Show article",
		"article": article,
	})
}

func (controller *Controller) Create(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{
		"title": "Create article",
	})
}

func (controller *Controller) Store(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "article created"})
}
