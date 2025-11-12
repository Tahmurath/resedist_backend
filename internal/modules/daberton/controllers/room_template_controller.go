package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	configStruct "resedist/config"

	//tgUserResponse "resedist/internal/modules/tgminiapp/responses"
	roomServices "resedist/internal/modules/daberton/services"
	//UserResponse "resedist/internal/modules/user/responses"
	authHelpers "resedist/internal/modules/auth/helpers"
	RoomRequest "resedist/internal/modules/daberton/requests"
	_ "resedist/internal/modules/daberton/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/rest"
)

type Controller struct {
	RoomService roomServices.RoomServiceInterface
	UserService UserService.UserServiceInterface
	errFmt      *errors.ErrorFormat
	json        *rest.Jsonresponse
	rest        configStruct.Rest
}

func New() *Controller {

	return &Controller{
		//departmentService: DepartmentService.New(),
		RoomService: roomServices.New(),
		UserService: UserService.New(),
		rest:        config.Get().Rest,
		errFmt:      errors.New(),
		json:        rest.New(),
	}
}

// @Summary Create room template
// @Description Returns room template json (requires JWT)
// @Security BearerAuth
// @ID create-room-template
// @Tags roomTemplate
// @Accept json
// @Produce json
// @Param request query RoomRequest.AddRoomTemplateRequest true "Department Create request"
// @Success 200 {object} _.RoomTemplate "Response object"
// @Router /api/v1/daberton/roomtemplate [post]
func (ctl *Controller) CreateRoomTemplate(c *gin.Context) {

	var request RoomRequest.AddRoomTemplateRequest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user := authHelpers.AuthJWT(c)
	_, err := ctl.RoomService.CreateRoomTemplate(request, user)
	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{
			Error_message: user,
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: request,
		Http: http.StatusCreated,
	})
}

// @Summary Create room template by admin
// @Description Returns room template json (requires JWT)
// @Security BearerAuth
// @ID admin-create-room-template
// @Tags AdminRoomTemplate
// @Accept json
// @Produce json
// @Param request query RoomRequest.AdminAddRoomTemplateRequest true "Department Create request"
// @Success 200 {object} _.RoomTemplate "Response object"
// @Router /api/v1/daberton/admin/roomtemplate [post]
func (ctl *Controller) AdminCreateRoomTemplate(c *gin.Context) {

	var request RoomRequest.AdminAddRoomTemplateRequest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user := authHelpers.AuthJWT(c)
	_, err := ctl.RoomService.AdminCreateRoomTemplate(request, user)
	if err != nil {
		ctl.json.ServerError(c, rest.RestConfig{
			Error_message: user,
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data: request,
		Http: http.StatusCreated,
	})
}

// @Summary Get room templates
// @Description Returns a list of room templates
// @Security BearerAuth
// @Tags roomTemplate
// @Accept json
// @Produce json
// @Param request query RoomRequest.AdminListRoomTemplateRequest true "Department Create request"
// @Success 200 {object} _.RoomTemplates"Response object"
// @Router /api/v1/daberton/roomtemplate [get]
func (ctl *Controller) AdminSearchRoomTemplate(c *gin.Context) {
	var request RoomRequest.AdminListRoomTemplateRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	rooms, paginate, err := ctl.RoomService.SearchRoomTemplatesPaginated(request)

	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	ctl.json.Success(c, rest.RestConfig{
		Data:       rooms.Data,
		Paged:      true,
		Pagination: paginate,
	})

}
