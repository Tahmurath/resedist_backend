package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	DepScopes "resedist/internal/modules/department/department/scopes"
	"resedist/pkg/config"
	"resedist/pkg/pagination"

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
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department, err := controller.departmentService.Find(
		request.DepartmentId,
		request.Expand,
		DepScopes.Preload(request.Expand, "DepartmentType", "Parent"),
	)

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
		DepScopes.IdsOr(request.Department),
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
			cfg.Error_message: err.Error(),  // دقیق‌تر کردن پیام خطا
			cfg.Error_code:    "BIND_ERROR", // یه کد خطای معنادار
		})
		return
	}

	user := authHelpers.AuthJWT(c)

	department, err := controller.departmentService.StoreAsUser(request, user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: err.Error(),
			cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		cfg.Status:        "success",
		cfg.Error_message: "",
		cfg.Error_code:    "",
		cfg.Data:          department,
	})
}

func (controller *Controller) Update(c *gin.Context) {
	var request DepRequest.OneDepartmentRequest

	cfg := config.Get().Jsonkey

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: err.Error(),
			cfg.Error_code:    "",
		})
		return
	}

	department, err := controller.departmentService.Find(request.DepartmentId, false)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: err.Error(),
			cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		cfg.Status:        "success",
		cfg.Error_message: "",
		cfg.Error_code:    "",
		cfg.Data:          department,
	})

}

//func (controller *Controller) Update(c *gin.Context) {
//	var request DepRequest.EditDepartmentRequest
//	cfg := config.Get().Jsonkey
//
//	id := c.Param("id")
//	if id == "" {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"request":         request,
//			cfg.Status:        "failed",
//			cfg.Error_message: "Department ID is required",
//			cfg.Error_code:    "INVALID_ID",
//		})
//		return
//	}
//
//	if err := c.ShouldBind(&request); err != nil {
//		c.JSON(http.StatusUnprocessableEntity, gin.H{
//			"request":         request,
//			cfg.Status:        "failed",
//			cfg.Error_message: err.Error(),
//			cfg.Error_code:    "BIND_ERROR",
//		})
//		return
//	}
//
//	user := authHelpers.AuthJWT(c)
//
//	department, err := controller.departmentService.UpdateAsUser(id, request, user)
//	if err != nil {
//		c.JSON(http.StatusUnprocessableEntity, gin.H{
//			"request":         request,
//			cfg.Status:        "failed",
//			cfg.Error_message: err.Error(),
//			cfg.Error_code:    "UPDATE_ERROR",
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		cfg.Status:        "success",
//		cfg.Error_message: "",
//		cfg.Error_code:    "",
//		cfg.Data:          department,
//	})
//}
