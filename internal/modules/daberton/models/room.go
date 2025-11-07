package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
	"time"
)

type RoomTemplate struct {
	gorm.Model
	Title      string  `gorm:"not null" json:"title"`
	EntryFee   float64 `gorm:"type:decimal(10,2);not null" json:"entry_fee"`
	MinPlayers int     `gorm:"not null" json:"min_players"`
	MaxPlayers int     `gorm:"not null" json:"max_players"`
	Timeout    int     `gorm:"not null" json:"timeout"`
	GameStyle  string  `gorm:"type:enum('tombola','Bingo','90-ball','azerbaijan','russia','iran','iraq','classic','modern','daberton');default:'classic'" json:"game_style"`
	IsPublic   bool    `gorm:"default:true;not null" json:"is_public"`
	CreatorID  int64   `gorm:"not null" json:"creator_id"`
	IsSystem   bool    `gorm:"default:false;not null" json:"is_system"`
	IsActive   bool    `gorm:"default:true;not null" json:"is_active"`

	Creator userModels.User `gorm:"foreignKey:CreatorID" json:"creator,omitempty"`
}

type RoomInstance struct {
	gorm.Model
	TemplateID   uint       `gorm:"not null" json:"template_id"`
	RoomStatus   string     `gorm:"type:enum('waiting','ready','in_progress','finished','cancelled');default:'waiting';not null" json:"room_status"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	FinishedAt   *time.Time `json:"finished_at,omitempty"`
	CancelReason *string    `json:"cancel_reason,omitempty"`

	Template RoomTemplate      `gorm:"foreignKey:TemplateID" json:"template"`
	Players  []userModels.User `gorm:"many2many:room_players;" json:"players"`
}

type RoomPlayer struct {
	gorm.Model
	InstanceID uint  `gorm:"primaryKey;not null" json:"instance_id"`
	UserID     int64 `gorm:"primaryKey;not null" json:"user_id"`
	HasPaid    bool  `gorm:"default:false;not null" json:"has_paid"`
	HasRefund  bool  `gorm:"default:false;not null" json:"has_refund"`

	User userModels.User `gorm:"foreignKey:UserID" json:"user"`
}

type RoomQueue struct {
	gorm.Model
	TemplateID uint  `gorm:"not null" json:"template_id"`
	UserID     int64 `gorm:"not null" json:"user_id"`

	Template RoomTemplate    `gorm:"foreignKey:TemplateID;not null" json:"-"`
	User     userModels.User `gorm:"foreignKey:UserID;not null" json:"user"`
}
