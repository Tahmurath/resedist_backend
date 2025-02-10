package models

import (
	"gorm.io/gorm"
	depTypeModel "resedist/internal/modules/department/department_type/models"
	userModels "resedist/internal/modules/user/models"
	//orderModels "resedist/internal/modules/order/models"
)

type Department struct {
	gorm.Model
	Title            string                       `gorm:"size:255;not null"`
	DepartmentTypeId uint                         `gorm:"index"`
	DepartmentType   *depTypeModel.DepartmentType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:DepartmentTypeId"`
	ParentID         uint                         `gorm:"index"`
	Parent           *Department                  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:ParentID"`
	Children         []Department                 `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AddedByUserID    uint
	AddedByUser      userModels.User
	//ModelHelper      ref.ModelHelper `gorm:"-"`
	//GetFieldName     ref.ModelHelper `gorm:"embedded"`
}

//DepartmentModel.Department{}.GetFieldName()
//departmentModel := DepartmentModel.Department{}

// استفاده از متد GetFieldName که از ModelHelper به ارث رسیده
//fmt.Println(departmentModel.ModelHelper.GetFieldName(departmentModel, "Title"))
