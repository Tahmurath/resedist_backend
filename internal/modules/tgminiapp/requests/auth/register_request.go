package auth

type TgRegisterRequest struct {
	TgID         int64  `json:"tg_id" binding:"required"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username" binding:"required"`
	LanguageCode string `json:"language_code" binding:"required"`
	PhotoURL     string `json:"photo_url"`
	IsBot        bool   `json:"is_bot"`
	IsPremium    bool   `json:"is_premium"`
}
