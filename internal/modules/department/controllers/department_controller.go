package controllers

import (
	"fmt"
	"net/http"
	"resedist/internal/modules/auth/helpers"
	"resedist/pkg/errors"

	"github.com/gin-gonic/gin"

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

func (controller *Controller) Search(c *gin.Context) {

	deptitle := c.DefaultQuery("query", "")
	departments := controller.departmentService.Search(deptitle)

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
