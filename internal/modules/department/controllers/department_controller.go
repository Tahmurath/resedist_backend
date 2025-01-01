package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/errors"
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

func (controller *Controller) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token":   "3453",
		"message": "User logged in successfully",
	})
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

	//user, _ := c.Get("auth")
	userService := UserService.New()
	user, _ := userService.GetCachedUserById(6)

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
