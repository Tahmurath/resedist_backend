package services

import (

	//UserRepository "resedist/internal/modules/user/repositories"
	//UserResponse "resedist/internal/modules/user/responses"
	"errors"
	roomRepository "resedist/internal/modules/daberton/repositories"

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

func (RoomService *RoomService) CreateRoomTemplate(request RoomRequest.RoomTemplateRequest, user UserResponse.User) (RoomResponse.RoomTemplate, error) {
	var response RoomResponse.RoomTemplate
	var room RoomModels.RoomTemplate

	room.CreatorID = int64(user.ID)
	room.EntryFee = request.EntryFee
	room.GameStyle = request.GameStyle
	room.IsPublic = *request.IsPublic
	room.MaxPlayers = request.MaxPlayers
	room.MinPlayers = request.MinPlayers
	room.Timeout = request.Timeout
	room.Title = request.Title
	room.IsActive = true
	room.IsSystem = false

	newRoom := RoomService.roomRepository.CreateTemplate(room)

	if newRoom.ID == 0 {
		return response, errors.New("user create fail")
	}

	return RoomResponse.ToRoomTemplate(newRoom), nil

}
