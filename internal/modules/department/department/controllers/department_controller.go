package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	DepScopes "resedist/internal/modules/department/department/scopes"
	"resedist/pkg/config"
	"resedist/pkg/pagination"
	"strconv"

	//articleRepository "resedist/internal/modules/article/repositories"

	authHelpers "resedist/internal/modules/auth/helpers"
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

func (controller *Controller) Show(c *gin.Context) {
	var request DepRequest.OneDepartmentRequest

	fmt.Println(strconv.Atoi(c.Param("expand")))

	cfg := config.Get().Jsonkey

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error with Query bind",
			cfg.Error_code:    "",
		})
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil { // گرفتن expand از JSON body
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department, err := controller.departmentService.Find(request.DepartmentId, request.Expand, DepScopes.Preload(request.Expand, "DepartmentType", "Parent"))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: "Department not found",
			cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		cfg.Status:        "",
		cfg.Error_message: "",
		cfg.Error_code:    "",
		cfg.Data:          department,
	})
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

	paginate := pagination.New(request.Page, request.PageSize)

	departments := controller.departmentService.SearchScope(
		request.Expand,
		paginate,
		DepScopes.TitleLike(request.Title),
		DepScopes.Preload(request.Expand, "DepartmentType", "Parent"),
		DepScopes.ParentIDS(request.Parent),
		DepScopes.DepTypes(request.DepartmentType),
		DepScopes.Sort(request.Sort, request.Order),
	)

	c.JSON(http.StatusOK, gin.H{
		cfg.Status:        "",
		cfg.Error_message: "",
		cfg.Error_code:    "",
		cfg.Pagination:    paginate,
		cfg.Data:          departments.Data,
	})
}

func (controller *Controller) Store(c *gin.Context) {
	var request DepRequest.AddDepartmentRequest

	cfg := config.Get().Jsonkey

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error with Query bind",
			cfg.Error_code:    "",
		})
		return
	}

	user := authHelpers.AuthJWT(c)

	department, err := controller.departmentService.StoreAsUser(request, user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error with store department",
			cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		cfg.Status:        "",
		cfg.Error_message: "",
		cfg.Error_code:    "",
		cfg.Data:          department,
	})
}
