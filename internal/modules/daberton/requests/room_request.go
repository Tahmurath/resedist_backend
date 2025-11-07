package requests

type RoomTemplateRequest struct {
	Title      string  `json:"title" binding:"required"`
	EntryFee   float64 `json:"entry_fee" binding:"required,gte=0"`
	MinPlayers int     `json:"min_players" binding:"required,gte=2"`
	MaxPlayers int     `json:"max_players" binding:"required,gtefield=MinPlayers"`
	Timeout    int     `json:"timeout" binding:"required,gt=0"`
	GameStyle  string  `json:"game_style" binding:"required,oneof=tombola Bingo 90-ball azerbaijan russia iran iraq classic modern daberton"`
	IsPublic   *bool   `json:"is_public" binding:"required"`
}

type CreateRoomInstanceRequest struct {
	TemplateID uint `json:"template_id" binding:"required"`
}

type AddPlayerToRoomRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type JoinQueueRequest struct {
	TemplateID uint `json:"template_id" binding:"required"`
}
