package responses

type TgRegisterResponse struct {
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
