package controllers

import (
	"net/http"
	configStruct "resedist/config"
	DepScopes "resedist/internal/modules/department/department/scopes"
	"resedist/pkg/config"
	"resedist/pkg/errors"

	"github.com/gin-gonic/gin"

	authHelpers "resedist/internal/modules/auth/helpers"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	DepartmentService "resedist/internal/modules/department/department/services"
)

type Controller struct {
	departmentService DepartmentService.DepartmentServiceInterface
	cfg               configStruct.Jsonkey
	error2            *errors.Error2
}

func New() *Controller {

	return &Controller{
		departmentService: DepartmentService.New(),
		cfg:               config.Get().Jsonkey,
		error2:            errors.New(),
	}
}

func (ctl *Controller) Show(c *gin.Context) {
	var request DepRequest.OneDepartmentRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":             request,
			ctl.cfg.Status:        "failed",
			ctl.cfg.Error_message: err.Error(),
			ctl.cfg.Error_code:    "",
		})
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department, err := ctl.departmentService.Find(
		request.DepartmentId,
		request.Expand,
		DepScopes.Preload(request.Expand, "DepartmentType", "Parent"),
	)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":             request,
			ctl.cfg.Status:        "failed",
			ctl.cfg.Error_message: "Department not found",
			ctl.cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.cfg.Status:        "",
		ctl.cfg.Error_message: "",
		ctl.cfg.Error_code:    "",
		ctl.cfg.Data:          department,
	})
}

func (ctl *Controller) Search(c *gin.Context) {
	var request DepRequest.ListDepartmentRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":             request,
			ctl.cfg.Status:        "failed",
			ctl.cfg.Error_message: "Opps, there is an error with Query bind",
			ctl.cfg.Error_code:    "",
		})
		return
	}

	departments, paginate, err := ctl.departmentService.SearchDepartmentsPaginated(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ctl.cfg.Status:        "failed",
			ctl.cfg.Error_message: "Error fetching departments",
			ctl.cfg.Error_code:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.cfg.Status:        "success",
		ctl.cfg.Error_message: "",
		ctl.cfg.Error_code:    "",
		ctl.cfg.Pagination:    paginate,
		ctl.cfg.Data:          departments.Data,
	})
}

func (ctl *Controller) Store(c *gin.Context) {
	var request DepRequest.AddDepartmentRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":             request,
			ctl.cfg.Status:        "failed",
			ctl.cfg.Error_message: err.Error(),
			ctl.cfg.Error_code:    "BIND_ERROR",
		})
		return
	}

	user := authHelpers.AuthJWT(c)

	department, err := ctl.departmentService.StoreAsUser(request, user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":             request,
			ctl.cfg.Status:        "failed",
			ctl.cfg.Error_message: err.Error(),
			ctl.cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.cfg.Status:        "success",
		ctl.cfg.Error_message: "",
		ctl.cfg.Error_code:    "",
		ctl.cfg.Data:          department,
	})
}

func (ctl *Controller) Update(c *gin.Context) {
	var request DepRequest.EditDepartmentRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":             request,
			ctl.cfg.Status:        "BIND_ERROR",
			ctl.cfg.Error_message: ctl.error2.SetFromError(err),
			ctl.cfg.Error_code:    "",
		})
		return
	}

	user := authHelpers.AuthJWT(c)
	department, err := ctl.departmentService.UpdateDepartment(request, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ctl.cfg.Status:        "error",
			ctl.cfg.Error_message: err.Error(),
			ctl.cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.cfg.Status:        "success",
		ctl.cfg.Error_message: "",
		ctl.cfg.Error_code:    "",
		ctl.cfg.Data:          department,
	})
}

func (ctl *Controller) Remove(c *gin.Context) {
	var request DepRequest.RemoveDepartmentRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":             request,
			ctl.cfg.Status:        "BIND_ERROR",
			ctl.cfg.Error_message: ctl.error2.SetFromError(err),
			ctl.cfg.Error_code:    "",
		})
		return
	}

	err := ctl.departmentService.Delete(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ctl.cfg.Status:        "error",
			ctl.cfg.Error_message: err.Error(),
			ctl.cfg.Error_code:    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.cfg.Status:        "success",
		ctl.cfg.Error_message: "",
		ctl.cfg.Error_code:    "",
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
