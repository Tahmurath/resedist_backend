package models

import (
	depTypeModel "resedist/internal/modules/department/department_type/models"
	userModels "resedist/internal/modules/user/models"

	"gorm.io/gorm"
	//orderModels "resedist/internal/modules/order/models"
)

type Department struct {
	gorm.Model
	Title            string                      `gorm:"size:255;not null"`
	DepartmentTypeId uint                        `gorm:"index"`
	DepartmentType   depTypeModel.DepartmentType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ParentID         uint                        `gorm:"index"`
	Parent           *Department                 `gorm:"foreignKey:ParentID"`
	Children         []Department                `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AddedByUserID    uint
	AddedByUser      userModels.User
}
