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
	DepTypeRequest "resedist/internal/modules/department/department_type/requests/deptype"
	_ "resedist/internal/modules/department/department_type/responses"
	DepTypeService "resedist/internal/modules/department/department_type/services"
)

type Controller struct {
	DepTypeService DepTypeService.DepTypeServiceInterface
	errFmt         *errors.ErrorFormat
	json           *rest.Jsonresponse
	rest           configStruct.Rest
}

func New() *Controller {

	return &Controller{
		DepTypeService: DepTypeService.New(),
		rest:           config.Get().Rest,
		errFmt:         errors.New(),
		json:           rest.New(),
	}
}

// @Summary Get Department type
// @Description Returns a Deaprtment type
// @Security BearerAuth
// @Tags depType
// @Accept json
// @Produce json
// @Param request path DepTypeRequest.OneDepTypeRequest true "Department type id"
// @Param request query DepTypeRequest.ShowDepTypeRequest true "Department type request"
// @Success 200 {object} _.DepTypeResponse "Response object"
// @Router /api/v1/department-type/{id} [get]
func (ctl *Controller) Show(c *gin.Context) {
	var uri DepTypeRequest.OneDepTypeRequest
	var request DepTypeRequest.ShowDepTypeRequest

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
	department, err := ctl.DepTypeService.Find(
		uri.DepTypeId,
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

// @Summary Get Department types
// @Description Returns a list of Deaprtment types
// @Security BearerAuth
// @Tags depType
// @Accept json
// @Produce json
// @Param user query DepTypeRequest.ListDepTypeRequest true "Department type search"
// @Success 200 {object} _.DepTypesResponse "Response object"
// @Router /api/v1/department-type/ [get]
func (ctl *Controller) Search(c *gin.Context) {
	var request DepTypeRequest.ListDepTypeRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	departments, paginate, err := ctl.DepTypeService.SearchDepTypesPaginated(request)

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

// @Summary create Department type
// @Description Returns a Department type (requires JWT)
// @Security BearerAuth
// @Tags depType
// @Accept json
// @Produce json
// @Param request query DepTypeRequest.AddDepTypeRequest true "Department type create request"
// @Success 200 {object} _.DepTypeResponse "Response object"
// @Router /api/v1/department-type/ [post]
func (ctl *Controller) Store(c *gin.Context) {
	var request DepTypeRequest.AddDepTypeRequest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user := authHelpers.AuthJWT(c)

	department, err := ctl.DepTypeService.StoreAsUser(request, user)
	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{
			Error_message: user,
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: department,
		Http: http.StatusCreated,
	})
}

// @Summary Update Department
// @Description Returns a Department type (requires JWT)
// @Security BearerAuth
// @Tags depType
// @Accept json
// @Produce json
// @Param request path DepTypeRequest.OneDepTypeRequest true "Department type id"
// @Param request query DepTypeRequest.EditDepTypeRequest true "Department type update request"
// @Success 200 {object} _.DepTypeResponse "Response object"
// @Router /api/v1/department-type/{id} [put]
func (ctl *Controller) Update(c *gin.Context) {
	var request DepTypeRequest.EditDepTypeRequest
	var uri DepTypeRequest.OneDepTypeRequest

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

	department, err := ctl.DepTypeService.UpdateDepartment(uri.DepTypeId, request, user)

	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: department,
	})
}

// @Summary Delete Department type
// @Description Returns No content (requires JWT)
// @Security BearerAuth
// @Tags depType
// @Accept json
// @Produce json
// @Param request path DepTypeRequest.OneDepTypeRequest true "Department type id"
// @Success 204 {object} _.NoContentResponse "Response object"
// @Router /api/v1/department-type/{id} [delete]
func (ctl *Controller) Remove(c *gin.Context) {
	var uri DepTypeRequest.OneDepTypeRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	err := ctl.DepTypeService.Delete(uri.DepTypeId)

	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Http:      http.StatusNoContent,
		NoContent: true,
	})
}
