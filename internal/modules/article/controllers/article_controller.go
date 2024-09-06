package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resedist/pkg/html"
	"strconv"

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
