package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
	//orderModels "resedist/internal/modules/order/models"
)

type Tenant struct {
	gorm.Model
	Title         string `gorm:"size:255;not null"`
	UserID        uint
	User          userModels.User
	AddedByUserID uint
	AddedByUser   userModels.User
	//ModelHelper      ref.ModelHelper `gorm:"-"`
	//GetFieldName     ref.ModelHelper `gorm:"embedded"`
}

//DepartmentModel.Department{}.GetFieldName()
//departmentModel := DepartmentModel.Department{}

// استفاده از متد GetFieldName که از ModelHelper به ارث رسیده
//fmt.Println(departmentModel.ModelHelper.GetFieldName(departmentModel, "Title"))
