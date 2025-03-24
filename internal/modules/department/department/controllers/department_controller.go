package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	configStruct "resedist/config"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/rest"

	authHelpers "resedist/internal/modules/auth/helpers"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	_ "resedist/internal/modules/department/department/responses"
	DepartmentService "resedist/internal/modules/department/department/services"
)

type Controller struct {
	departmentService DepartmentService.DepartmentServiceInterface
	errFmt            *errors.ErrorFormat
	json              *rest.Jsonresponse
	rest              configStruct.Rest
}

func New() *Controller {

	return &Controller{
		departmentService: DepartmentService.New(),
		rest:              config.Get().Rest,
		errFmt:            errors.New(),
		json:              rest.New(),
	}
}

// @Summary Get Department
// @Description Returns a Deaprtment
// @Security BearerAuth
// @Tags department
// @Accept json
// @Produce json
// @Param request path DepRequest.OneDepartmentRequest true "Department request data"
// @Param request query DepRequest.ShowDepartmentRequest true "Department request data"
// @Success 200 {object} _.DepartmentResponse "Response object"
// @Router /api/v1/department/{id} [get]
func (ctl *Controller) Show(c *gin.Context) {
	var uri DepRequest.OneDepartmentRequest
	var request DepRequest.ShowDepartmentRequest

	if err := c.ShouldBindUri(&uri); err != nil {
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
		uri.DepartmentId,
		request.Expand,
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

	// c.JSON(http.StatusOK, DepResponse.DepartmentResponse{
	// 	ErrorCode: "",
	// 	Status:    ctl.rest.Success,
	// 	Data:      department,
	// 	Message:   "",
	// })

}

// @Summary Get Departments
// @Description Returns a list ofDeaprtment
// @Security BearerAuth
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

	// c.JSON(http.StatusOK, DepResponse.DepartmentsResponse{
	// 	ErrorCode:  "",
	// 	Status:     ctl.rest.Success,
	// 	Data:       departments.Data,
	// 	Message:    "",
	// 	Pagination: paginate,
	// })
}

// @Summary Get Department
// @Description Returns a Department (requires JWT)
// @Security BearerAuth
// @Tags department
// @Accept json
// @Produce json
// @Param request query DepRequest.AddDepartmentRequest true "Department request data"
// @Success 200 {object} _.DepartmentResponse "Response object"
// @Router /api/v1/department/ [post]
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

// @Summary Update Department
// @Description Returns a Department (requires JWT)
// @Security BearerAuth
// @Tags department
// @Accept json
// @Produce json
// @Param request path DepRequest.OneDepartmentRequest true "Department request data"
// @Param request query DepRequest.EditDepartmentRequest true "Department request data"
// @Success 200 {object} _.DepartmentResponse "Response object"
// @Router /api/v1/department/{id} [put]
func (ctl *Controller) Update(c *gin.Context) {
	var request DepRequest.EditDepartmentRequest
	var uri DepRequest.OneDepartmentRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		log.Printf("ShouldBindUri: %+v\n", err)
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	if err := c.ShouldBind(&request); err != nil {
		log.Printf("ShouldBind: %+v\n", request)
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user := authHelpers.AuthJWT(c)

	department, err := ctl.departmentService.UpdateDepartment(uri.DepartmentId, request, user)

	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: department,
	})

	// c.JSON(http.StatusOK, DepResponse.DepartmentResponse{
	// 	ErrorCode: "",
	// 	Status:    ctl.rest.Success,
	// 	Data:      department,
	// 	Message:   "",
	// })
}

// @Summary Delete Department
// @Description Returns No content (requires JWT)
// @Security BearerAuth
// @Tags department
// @Accept json
// @Produce json
// @Param request path DepRequest.OneDepartmentRequest true "Department request data"
// @Success 204 {object} _.NoContentResponse "Response object"
// @Router /api/v1/department/{id} [delete]
func (ctl *Controller) Remove(c *gin.Context) {
	var uri DepRequest.OneDepartmentRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	err := ctl.departmentService.Delete(uri.DepartmentId)

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
