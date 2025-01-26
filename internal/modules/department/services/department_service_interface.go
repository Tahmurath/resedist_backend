package services

import (
	"gorm.io/gorm"
	DepRequest "resedist/internal/modules/department/requests/department"
	DepResponse "resedist/internal/modules/department/responses"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"
)

type DepartmentServiceInterface interface {
	//Find(id int) (DepResponse.Department, error)
	//List() DepResponse.Departments
	StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)
	Search(title string, page int, pageSize int, expand bool) DepResponse.Departments
	SearchP(pack *pagination.PagePack) DepResponse.Departments
	SearchScope(expand bool, pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) DepResponse.Departments
}
