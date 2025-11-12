package services

import (

	//UserRepository "resedist/internal/modules/user/repositories"
	//UserResponse "resedist/internal/modules/user/responses"
	"errors"
	"gorm.io/gorm"
	roomRepository "resedist/internal/modules/daberton/repositories"
	RoomTemplateScopes "resedist/internal/modules/daberton/scopes"
	"resedist/pkg/pagination"

	//UserResponse "resedist/internal/modules/user/responses"
	RoomModels "resedist/internal/modules/daberton/models"
	RoomRequest "resedist/internal/modules/daberton/requests"
	RoomResponse "resedist/internal/modules/daberton/responses"
	UserResponse "resedist/internal/modules/user/responses"
)

type RoomService struct {
	roomRepository roomRepository.RoomRepositoryInterface
}

func New() *RoomService {
	return &RoomService{
		roomRepository: roomRepository.New(),
	}
}

func (RoomService *RoomService) CreateRoomTemplate(request RoomRequest.AddRoomTemplateRequest, user UserResponse.User) (RoomResponse.RoomTemplate, error) {
	var response RoomResponse.RoomTemplate
	var room RoomModels.RoomTemplate

	room.CreatorID = int64(user.ID)
	room.EntryFee = request.EntryFee
	room.GameStyle = request.GameStyle
	room.MaxPlayers = request.MaxPlayers
	room.MinPlayers = request.MinPlayers
	room.Timeout = request.Timeout
	room.Title = request.Title
	room.IsPublic = *request.IsPublic
	room.IsActive = *request.IsActive
	room.IsSystem = false

	newRoom := RoomService.roomRepository.CreateTemplate(room)

	if newRoom.ID == 0 {
		return response, errors.New("room template create fail")
	}

	return RoomResponse.ToRoomTemplate(newRoom), nil

}

func (RoomService *RoomService) AdminCreateRoomTemplate(request RoomRequest.AdminAddRoomTemplateRequest, user UserResponse.User) (RoomResponse.RoomTemplate, error) {
	var response RoomResponse.RoomTemplate
	var room RoomModels.RoomTemplate

	room.CreatorID = int64(user.ID)
	room.EntryFee = request.EntryFee
	room.GameStyle = request.GameStyle
	room.MaxPlayers = request.MaxPlayers
	room.MinPlayers = request.MinPlayers
	room.Timeout = request.Timeout
	room.Title = request.Title
	room.IsPublic = *request.IsPublic
	room.IsActive = *request.IsActive
	room.IsSystem = true

	newRoom := RoomService.roomRepository.CreateTemplate(room)

	if newRoom.ID == 0 {
		return response, errors.New("room template create fail")
	}

	return RoomResponse.ToRoomTemplate(newRoom), nil

}

func (RoomService *RoomService) SearchRoomTemplatesWithScopes(expand bool, pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) (RoomResponse.RoomTemplates, error) {
	rooms, err := RoomService.roomRepository.FindAllScope(pack, scopes...)
	if err != nil {
		return RoomResponse.RoomTemplates{}, err
	}
	return RoomResponse.ToRoomTemplates(rooms, expand), nil
}

func (RoomService *RoomService) SearchRoomTemplatesPaginated(request RoomRequest.AdminListRoomTemplateRequest) (RoomResponse.RoomTemplates, pagination.PagePack, error) {
	paginate := pagination.NewPagePack(request.Page, request.PageSize)

	scopes := []func(*gorm.DB) *gorm.DB{
		RoomTemplateScopes.TitleLike(request.Title),
		RoomTemplateScopes.Sort(request.Sort, request.Order),
	}

	rooms, err := RoomService.roomRepository.FindAllScope(paginate, scopes...)
	if err != nil {
		return RoomResponse.RoomTemplates{}, *paginate, err
	}
	return RoomResponse.ToRoomTemplates(rooms, request.Expand), *paginate, nil
}
