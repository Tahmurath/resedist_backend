package services

import (
	DepRequest "resedist/internal/modules/department/department/requests/department"
	DepResponse "resedist/internal/modules/department/department/responses"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentServiceInterface interface {
	Find(id uint, expand bool) (DepResponse.Department, error)
	StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)
	UpdateDepartment(id uint, request DepRequest.EditDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)
	Delete(id uint) error
	//UpdateAsUser(id int, request DepRequest.EditDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)

	SearchDepartmentsWithScopes(expand bool, pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) (DepResponse.Departments, error)
	SearchDepartmentsPaginated(request DepRequest.ListDepartmentRequest) (DepResponse.Departments, pagination.PagePack, error)
}
