package responses

import (
	roomModels "resedist/internal/modules/daberton/models"
	userResponses "resedist/internal/modules/user/responses"
	"time"
)

type RoomPlayer struct {
	ID         uint               `json:"id"`
	InstanceID uint               `json:"instance_id"`
	UserID     int64              `json:"user_id"`
	HasPaid    bool               `json:"has_paid"`
	HasRefund  bool               `json:"has_refund"`
	User       userResponses.User `json:"user"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type RoomPlayers struct {
	Data []RoomPlayer `json:"data"`
}

func ToRoomPlayer(roomPlayer roomModels.RoomPlayer) RoomPlayer {
	return RoomPlayer{
		ID:         roomPlayer.ID,
		InstanceID: roomPlayer.InstanceID,
		UserID:     roomPlayer.UserID,
		HasPaid:    roomPlayer.HasPaid,
		HasRefund:  roomPlayer.HasRefund,
		User:       userResponses.ToUser(roomPlayer.User),
		CreatedAt:  roomPlayer.CreatedAt,
		UpdatedAt:  roomPlayer.UpdatedAt,
	}
}

func ToRoomPlayers(roomPlayers []roomModels.RoomPlayer) RoomPlayers {
	var responseRoomPlayers []RoomPlayer
	for _, roomPlayer := range roomPlayers {
		responseRoomPlayers = append(responseRoomPlayers, ToRoomPlayer(roomPlayer))
	}
	return RoomPlayers{
		Data: responseRoomPlayers,
	}
}
