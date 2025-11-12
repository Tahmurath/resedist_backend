package requests

type AddRoomTemplateRequest struct {
	Title      string  `form:"title" json:"title" binding:"required"`
	EntryFee   float64 `form:"entry_fee" json:"entry_fee" binding:"required,gte=0"`
	MinPlayers int     `form:"min_players" json:"min_players" binding:"required,gte=2"`
	MaxPlayers int     `form:"max_players" json:"max_players" binding:"required,gtefield=MinPlayers"`
	Timeout    int     `form:"timeout" json:"timeout" binding:"required,gt=0"`
	GameStyle  string  `form:"game_style" json:"game_style" binding:"required,oneof=tombola Bingo 90-ball azerbaijan russia iran iraq classic modern daberton"`
	IsPublic   *bool   `form:"is_public" json:"is_public" binding:"required"`
	IsActive   *bool   `form:"is_active" json:"is_active"`
}

type AdminAddRoomTemplateRequest struct {
	Title      string  `form:"title" json:"title" binding:"required"`
	EntryFee   float64 `form:"entry_fee" json:"entry_fee" binding:"required,gte=0"`
	MinPlayers int     `form:"min_players" json:"min_players" binding:"required,gte=2"`
	MaxPlayers int     `form:"max_players" json:"max_players" binding:"required,gtefield=MinPlayers"`
	Timeout    int     `form:"timeout" json:"timeout" binding:"required,gt=0"`
	GameStyle  string  `form:"game_style" json:"game_style" binding:"required,oneof=tombola Bingo 90-ball azerbaijan russia iran iraq classic modern daberton"`
	IsPublic   *bool   `form:"is_public" json:"is_public" binding:"required"`
	IsActive   *bool   `form:"is_active" json:"is_active"`
}

type AdminListRoomTemplateRequest struct {
	Title    string  `form:"title" json:"title"`
	EntryFee float64 `form:"entry_fee" json:"entry_fee" binding:"gte=0"`
	//GameStyle string  `form:"game_style" json:"game_style" binding:"oneof=tombola Bingo 90-ball azerbaijan russia iran iraq classic modern daberton"`
	GameStyle string `form:"game_style" json:"game_style"`

	IsPublic  *bool  `form:"is_public" json:"is_public"`
	IsActive  *bool  `form:"is_active" json:"is_active"`
	IsSystem  *bool  `form:"is_system" json:"is_system"`
	CreatorID int64  `form:"creator_id" json:"creator_id"`
	Expand    bool   `form:"expand" json:"expand"`
	Sort      string `form:"sort" json:"sort"`
	Order     string `form:"order" json:"order"`
	Page      int    `form:"page" json:"page"`
	PageSize  int    `form:"page_size" json:"page_size"`
}

type ShowRoomTemplateRequest struct {
	Expand bool `form:"expand"`
}
