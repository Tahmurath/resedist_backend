package controllers

import (
	"net/http"
	configStruct "resedist/config"
	DepScopes "resedist/internal/modules/department/department/scopes"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/rest"

	"github.com/gin-gonic/gin"

	authHelpers "resedist/internal/modules/auth/helpers"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	DepartmentService "resedist/internal/modules/department/department/services"
)

type Controller struct {
	departmentService DepartmentService.DepartmentServiceInterface
	rest              configStruct.Rest
	errFmt            *errors.ErrorFormat
	json              *rest.Jsonresponse
}

func New() *Controller {

	return &Controller{
		departmentService: DepartmentService.New(),
		rest:              config.Get().Rest,
		errFmt:            errors.New(),
		json:              rest.New(),
	}
}

func (ctl *Controller) Show(c *gin.Context) {
	var request DepRequest.OneDepartmentRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: ctl.errFmt.SetFromError(err),
			ctl.rest.Error_code:    ctl.rest.Bind_error,
		})
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: ctl.errFmt.SetFromError(err),
			ctl.rest.Error_code:    ctl.rest.Bind_error,
		})
		return
	}
	department, err := ctl.departmentService.Find(
		request.DepartmentId,
		request.Expand,
		DepScopes.Preload(request.Expand, "DepartmentType", "Parent"),
	)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: err.Error(),
			ctl.rest.Error_code:    ctl.rest.Not_found,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.rest.Status:        ctl.rest.Success,
		ctl.rest.Error_message: "",
		ctl.rest.Error_code:    "",
		ctl.rest.Data:          department,
	})
}

// @Summary Get Departments
// @Description Returns a list ofDeaprtment
// @Tags department
// @Accept json
// @Produce json
// @Param user query DepRequest.ListDepartmentRequest true "User data"
// @Success 200 {object} DepRequest.ListDepartmentRequest "Response object"
// @Router /api/v1/department/ [get]
func (ctl *Controller) Search(c *gin.Context) {
	var request DepRequest.ListDepartmentRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: ctl.errFmt.SetFromError(err),
			ctl.rest.Error_code:    ctl.rest.Bind_error,
		})
		return
	}

	departments, paginate, err := ctl.departmentService.SearchDepartmentsPaginated(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: err.Error(),
			ctl.rest.Error_code:    ctl.rest.Not_found,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.rest.Status:        ctl.rest.Success,
		ctl.rest.Error_message: "",
		ctl.rest.Error_code:    "",
		ctl.rest.Pagination:    paginate,
		ctl.rest.Data:          departments.Data,
	})
}

func (ctl *Controller) Store(c *gin.Context) {
	var request DepRequest.AddDepartmentRequest

	if err := c.ShouldBindUri(&request); err != nil {

		ctl.json.Badrequest(c, rest.BadrequestConfig{
			Error_code:    "custome",
			Error_message: ctl.errFmt.SetFromError(err),
		})

		// c.JSON(http.StatusBadRequest, gin.H{
		// 	ctl.rest.Status:        ctl.rest.Failed,
		// 	ctl.rest.Error_message: ctl.errFmt.SetFromError(err),
		// 	ctl.rest.Error_code:    ctl.rest.Bind_error,
		// })
		return
	}

	user := authHelpers.AuthJWT(c)

	department, err := ctl.departmentService.StoreAsUser(request, user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: err.Error(),
			ctl.rest.Error_code:    ctl.rest.Not_found,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		ctl.rest.Status:        ctl.rest.Success,
		ctl.rest.Error_message: "",
		ctl.rest.Error_code:    "",
		ctl.rest.Data:          department,
	})
}

func (ctl *Controller) Update(c *gin.Context) {
	var request DepRequest.EditDepartmentRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: ctl.errFmt.SetFromError(err),
			ctl.rest.Error_code:    ctl.rest.Bind_error,
		})
		return
	}

	user := authHelpers.AuthJWT(c)
	department, err := ctl.departmentService.UpdateDepartment(request, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: err.Error(),
			ctl.rest.Error_code:    ctl.rest.Not_found,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.rest.Status:        ctl.rest.Success,
		ctl.rest.Error_message: "",
		ctl.rest.Error_code:    "",
		ctl.rest.Data:          department,
	})
}

func (ctl *Controller) Remove(c *gin.Context) {
	var request DepRequest.RemoveDepartmentRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: ctl.errFmt.SetFromError(err),
			ctl.rest.Error_code:    ctl.rest.Bind_error,
		})
		return
	}

	err := ctl.departmentService.Delete(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ctl.rest.Status:        ctl.rest.Failed,
			ctl.rest.Error_message: err.Error(),
			ctl.rest.Error_code:    ctl.rest.Not_found,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		ctl.rest.Status:        ctl.rest.Success,
		ctl.rest.Error_message: "",
		ctl.rest.Error_code:    "",
	})
}
