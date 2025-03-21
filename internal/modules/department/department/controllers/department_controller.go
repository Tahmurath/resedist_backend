package controllers

import (
	"net/http"
	DepScopes "resedist/internal/modules/department/department/scopes"
	"resedist/pkg/errors"
	"resedist/pkg/rest"

	"github.com/gin-gonic/gin"

	authHelpers "resedist/internal/modules/auth/helpers"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	_ "resedist/internal/modules/department/department/responses"
	DepartmentService "resedist/internal/modules/department/department/services"
)

type Controller struct {
	departmentService DepartmentService.DepartmentServiceInterface
	errFmt            *errors.ErrorFormat
	json              *rest.Jsonresponse
}

func New() *Controller {

	return &Controller{
		departmentService: DepartmentService.New(),
		errFmt:            errors.New(),
		json:              rest.New(),
	}
}

// @Summary Get Department
// @Description Returns a Deaprtment
// @Tags department
// @Accept json
// @Produce json
// @Param request path DepRequest.OneDepartmentRequest true "Department request data"
// @Param request query DepRequest.OneDepartmentRequest true "Department request data"
// @Success 200 {object} _.DepartmentResponse "Response object"
// @Router /api/v1/department/{id} [get]
func (ctl *Controller) Show(c *gin.Context) {
	var request DepRequest.OneDepartmentRequest

	if err := c.ShouldBindUri(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}
	department, err := ctl.departmentService.Find(
		request.DepartmentId,
		request.Expand,
		DepScopes.Preload(request.Expand, "DepartmentType", "Parent"),
	)

	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: department,
	})

}

// @Summary Get Departments
// @Description Returns a list ofDeaprtment
// @Tags department
// @Accept json
// @Produce json
// @Param user query DepRequest.ListDepartmentRequest true "User data"
// @Success 200 {object} _.DepartmentsResponse "Response object"
// @Router /api/v1/department/ [get]
func (ctl *Controller) Search(c *gin.Context) {
	var request DepRequest.ListDepartmentRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	departments, paginate, err := ctl.departmentService.SearchDepartmentsPaginated(request)

	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data:       departments.Data,
		Paged:      true,
		Pagination: paginate,
	})
}

func (ctl *Controller) Store(c *gin.Context) {
	var request DepRequest.AddDepartmentRequest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user := authHelpers.AuthJWT(c)

	department, err := ctl.departmentService.StoreAsUser(request, user)
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: department,
		Http: http.StatusCreated,
	})
}

func (ctl *Controller) Update(c *gin.Context) {
	var request DepRequest.EditDepartmentRequest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user := authHelpers.AuthJWT(c)
	department, err := ctl.departmentService.UpdateDepartment(request, user)

	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: department,
	})
}

func (ctl *Controller) Remove(c *gin.Context) {
	var request DepRequest.RemoveDepartmentRequest

	if err := c.ShouldBindUri(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	err := ctl.departmentService.Delete(request)

	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Http:      http.StatusNoContent,
		NoContent: true,
	})
}
