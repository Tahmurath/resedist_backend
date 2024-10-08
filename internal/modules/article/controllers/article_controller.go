package controllers

import (
	"fmt"
	"net/http"
	"resedist/pkg/converters"
	"resedist/pkg/errors"
	"resedist/pkg/html"
	"resedist/pkg/old"
	"resedist/pkg/sessions"
	"strconv"

	"github.com/gin-gonic/gin"

	//articleRepository "resedist/internal/modules/article/repositories"

	"resedist/internal/modules/article/requests/articles"
	ArticleService "resedist/internal/modules/article/services"
	"resedist/internal/modules/user/helpers"
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

	var request articles.StoreRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {

		errors.Init()
		errors.SetFromError(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	// create article
	article, err := controller.articleService.StoreAsUser(request, user)
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
