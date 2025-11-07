package responses

import (
	roomModels "resedist/internal/modules/daberton/models"
	userModels "resedist/internal/modules/user/models"
	userResponses "resedist/internal/modules/user/responses"
	"time"
)

type RoomInstance struct {
	ID           uint                 `json:"id"`
	Template     RoomTemplate         `json:"template"`
	RoomStatus   string               `json:"room_status"`
	StartedAt    *time.Time           `json:"started_at,omitempty"`
	FinishedAt   *time.Time           `json:"finished_at,omitempty"`
	CancelReason *string              `json:"cancel_reason,omitempty"`
	Players      []userResponses.User `json:"players"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
}

type RoomInstances struct {
	Data []RoomInstance `json:"data"`
}

// toUserResponses converts a slice of user models to a slice of user responses.
func toUserResponses(players []userModels.User) []userResponses.User {
	var userResps []userResponses.User
	for _, player := range players {
		userResps = append(userResps, userResponses.ToUser(player))
	}
	return userResps
}

func ToRoomInstance(roomInstance roomModels.RoomInstance) RoomInstance {
	return RoomInstance{
		ID:           roomInstance.ID,
		Template:     ToRoomTemplate(roomInstance.Template),
		RoomStatus:   roomInstance.RoomStatus,
		StartedAt:    roomInstance.StartedAt,
		FinishedAt:   roomInstance.FinishedAt,
		CancelReason: roomInstance.CancelReason,
		Players:      toUserResponses(roomInstance.Players), // اصلاح شده
		CreatedAt:    roomInstance.CreatedAt,
		UpdatedAt:    roomInstance.UpdatedAt,
	}
}

func ToRoomInstances(roomInstances []roomModels.RoomInstance) RoomInstances {
	var responseRoomInstances []RoomInstance
	for _, roomInstance := range roomInstances {
		responseRoomInstances = append(responseRoomInstances, ToRoomInstance(roomInstance))
	}
	return RoomInstances{
		Data: responseRoomInstances,
	}
}
