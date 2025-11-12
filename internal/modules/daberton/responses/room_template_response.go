package responses

import (
	roomModels "resedist/internal/modules/daberton/models"
	userResponses "resedist/internal/modules/user/responses"
	"time"
)

type RoomTemplate struct {
	ID         uint               `json:"id"`
	Title      string             `json:"title"`
	EntryFee   float64            `json:"entry_fee"`
	MinPlayers int                `json:"min_players"`
	MaxPlayers int                `json:"max_players"`
	Timeout    int                `json:"timeout"`
	GameStyle  string             `json:"game_style"`
	IsPublic   bool               `json:"is_public"`
	IsSystem   bool               `json:"is_system"`
	IsActive   bool               `json:"is_active"`
	Creator    userResponses.User `json:"creator,omitempty"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type RoomTemplates struct {
	Data []RoomTemplate `json:"data"`
}

func ToRoomTemplate(roomTemplate roomModels.RoomTemplate) RoomTemplate {
	return RoomTemplate{
		ID:         roomTemplate.ID,
		Title:      roomTemplate.Title,
		EntryFee:   roomTemplate.EntryFee,
		MinPlayers: roomTemplate.MinPlayers,
		MaxPlayers: roomTemplate.MaxPlayers,
		Timeout:    roomTemplate.Timeout,
		GameStyle:  roomTemplate.GameStyle,
		IsPublic:   roomTemplate.IsPublic,
		IsSystem:   roomTemplate.IsSystem,
		IsActive:   roomTemplate.IsActive,
		Creator:    userResponses.ToUser(roomTemplate.Creator),
		CreatedAt:  roomTemplate.CreatedAt,
		UpdatedAt:  roomTemplate.UpdatedAt,
	}
}

func ToRoomTemplates(roomTemplates []roomModels.RoomTemplate, expand bool) RoomTemplates {
	var responseRoomTemplates []RoomTemplate

	for _, roomTemplate := range roomTemplates {
		responseRoomTemplates = append(responseRoomTemplates, ToRoomTemplate(roomTemplate))
	}

	return RoomTemplates{
		Data: responseRoomTemplates,
	}
}
