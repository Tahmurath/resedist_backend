package requests

type RoomTemplateRequest struct {
	Title      string  `form:"title" json:"title" binding:"required"`
	EntryFee   float64 `form:"entry_fee" json:"entry_fee" binding:"required,gte=0"`
	MinPlayers int     `form:"min_players" json:"min_players" binding:"required,gte=2"`
	MaxPlayers int     `form:"max_players" json:"max_players" binding:"required,gtefield=MinPlayers"`
	Timeout    int     `form:"timeout" json:"timeout" binding:"required,gt=0"`
	GameStyle  string  `form:"game_style" json:"game_style" binding:"required,oneof=tombola Bingo 90-ball azerbaijan russia iran iraq classic modern daberton"`
	IsPublic   *bool   `form:"is_public" json:"is_public" binding:"required"`
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
