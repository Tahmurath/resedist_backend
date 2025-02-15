package controllers

import (
	"fmt"
	"net/http"
	"resedist/internal/modules/auth/helpers"
	DepScopes "resedist/internal/modules/department/department/scopes"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/pagination"

	"github.com/gin-gonic/gin"

	//articleRepository "resedist/internal/modules/article/repositories"

	DepRequest "resedist/internal/modules/department/department/requests/department"
	DepartmentService "resedist/internal/modules/department/department/services"
)

type Controller struct {
	departmentService DepartmentService.DepartmentServiceInterface
}

func New() *Controller {

	return &Controller{
		departmentService: DepartmentService.New(),
	}
}

func (controller *Controller) Search(c *gin.Context) {
	var request DepRequest.ListDepartmentRequest

	cfg := config.Get().Jsonkey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error with Query bind",
			cfg.Error_code:    "",
		})
		return
	}

	page := pagination.New(request.Page, request.PageSize)

	departments := controller.departmentService.SearchScope(
		request.Expand,
		page,
		DepScopes.TitleLike(request.Title),
		DepScopes.Preload(request.Expand, "DepartmentType", "Parent"),
		DepScopes.ParentID(request.ParentID),
		DepScopes.ParentIDS(request.Parent),
		DepScopes.DepTypes(request.DepartmentType),
		DepScopes.Sort(request.Sort, request.Order),
	)

	c.JSON(http.StatusOK, gin.H{
		cfg.Status:        "",
		cfg.Error_message: "",
		cfg.Error_code:    "",
		cfg.Pagination:    page,
		cfg.Data:          departments.Data,
	})
}

//func (controller *Controller) Search(c *gin.Context) {
//
//	title := c.DefaultQuery("query", "")
//	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
//	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
//	expand := c.Query("expand") == "true"
//
//	departments := controller.departmentService.Search(title, page, pageSize, expand)
//
//	c.JSON(http.StatusOK, departments.Data)
//}

func (controller *Controller) Store(c *gin.Context) {

	var request DepRequest.AddDepartmentRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {

		fmt.Println(err)
		errors.Init()
		errors.SetFromError(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with ShouldBind",
			"errors":  errors.Get(),
			"request": request,
		})
		return
	}

	user := helpers.AuthJWT(c)

	// create article
	department, err := controller.departmentService.StoreAsUser(request, user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with ShouldBind",
			"errors":  errors.Get(),
			"request": request,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Department createdd in successfully",
		"department": department,
	})
}
