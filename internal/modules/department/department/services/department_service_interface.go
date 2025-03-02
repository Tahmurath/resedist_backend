package services

import (
	"gorm.io/gorm"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	DepResponse "resedist/internal/modules/department/department/responses"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"
)

type DepartmentServiceInterface interface {
	Find(id int, expand bool, scopes ...func(*gorm.DB) *gorm.DB) (DepResponse.Department, error)
	StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)
	SearchScope(expand bool, pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) DepResponse.Departments
}
