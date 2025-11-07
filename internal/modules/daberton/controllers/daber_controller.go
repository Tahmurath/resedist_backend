package controllers

import (
	"net/http"
	configStruct "resedist/config"

	"github.com/gin-gonic/gin"

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
// @Param request query RoomRequest.RoomTemplateRequest true "Department Create request"
// @Success 200 {object} _.RoomTemplate "Response object"
// @Router /api/v1/daberton/roomtemplate [post]
func (ctl *Controller) CreateRoomTemplate(c *gin.Context) {

	var request RoomRequest.RoomTemplateRequest

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
