package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}
