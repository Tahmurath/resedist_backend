package responses

import (
	tgUserModels "resedist/internal/modules/tgminiapp/models"
)

type TgUser struct {
	ID           uint   `json:"id"`
	TgID         int64  `json:"tg_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	PhotoURL     string `json:"photo_url"`
	IsBot        bool   `json:"is_bot"`
	IsPremium    bool   `json:"is_premium"`
	UserID       uint   `json:"user_id"`
}

type TgUsers struct {
	Data []TgUser
}

func ToTgUser(tgUser tgUserModels.TgUser) TgUser {

	return TgUser{
		ID:           tgUser.ID,
		TgID:         tgUser.TgID,
		UserID:       tgUser.UserID,
		FirstName:    tgUser.FirstName,
		LastName:     tgUser.LastName,
		Username:     tgUser.Username,
		LanguageCode: tgUser.LanguageCode,
		PhotoURL:     tgUser.PhotoURL,
		IsBot:        tgUser.IsBot,
		IsPremium:    tgUser.IsPremium,
	}
}
func ToTgUsers(tgUsers []tgUserModels.TgUser) TgUsers {
	var responseTgUsers []TgUser

	for _, tgUser := range tgUsers {
		responseTgUsers = append(responseTgUsers, ToTgUser(tgUser))
	}

	return TgUsers{
		Data: responseTgUsers,
	}
}
