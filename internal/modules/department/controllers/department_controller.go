package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resedist/internal/modules/auth/helpers"
	"resedist/pkg/errors"
	"resedist/pkg/pagination"
	"strconv"

	//articleRepository "resedist/internal/modules/article/repositories"

	DepRequest "resedist/internal/modules/department/requests/department"
	DepartmentService "resedist/internal/modules/department/services"
)

type Controller struct {
	departmentService DepartmentService.DepartmentServiceInterface
}

func New() *Controller {

	return &Controller{
		departmentService: DepartmentService.New(),
	}
}

func (controller *Controller) Search2(c *gin.Context) {

	//pp := pagination.NewConfig(c, "page", "page_size", "expand", "search")
	pp := pagination.New(c)

	departments := controller.departmentService.SearchP(pp)

	c.JSON(http.StatusOK, gin.H{
		"status":        "",
		"error_message": "",
		"error_code":    "",
		"_metadata":     pp,
		"data":          departments.Data,
	})
}
func (controller *Controller) Search(c *gin.Context) {

	title := c.DefaultQuery("query", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	expand := c.Query("expand") == "true"

	departments := controller.departmentService.Search(title, page, pageSize, expand)

	c.JSON(http.StatusOK, departments.Data)
}

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
