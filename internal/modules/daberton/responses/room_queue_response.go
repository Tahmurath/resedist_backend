package responses

import (
	roomModels "resedist/internal/modules/daberton/models"
	userResponses "resedist/internal/modules/user/responses"
	"time"
)

type RoomQueue struct {
	ID         uint               `json:"id"`
	TemplateID uint               `json:"template_id"`
	UserID     int64              `json:"user_id"`
	User       userResponses.User `json:"user"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type RoomQueues struct {
	Data []RoomQueue `json:"data"`
}

func ToRoomQueue(roomQueue roomModels.RoomQueue) RoomQueue {
	return RoomQueue{
		ID:         roomQueue.ID,
		TemplateID: roomQueue.TemplateID,
		UserID:     roomQueue.UserID,
		User:       userResponses.ToUser(roomQueue.User),
		CreatedAt:  roomQueue.CreatedAt,
		UpdatedAt:  roomQueue.UpdatedAt,
	}
}

func ToRoomQueues(roomQueues []roomModels.RoomQueue) RoomQueues {
	var responseRoomQueues []RoomQueue
	for _, roomQueue := range roomQueues {
		responseRoomQueues = append(responseRoomQueues, ToRoomQueue(roomQueue))
	}
	return RoomQueues{
		Data: responseRoomQueues,
	}
}
