package services

import (
	DepTypeResponse "resedist/internal/modules/department/department_type/responses"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentTypeServiceInterface interface {
	//Find(id int) (DepResponse.Department, error)
	Search(title string) DepTypeResponse.DepartmentTypes
	SearchScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) DepTypeResponse.DepartmentTypes
	// StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)
}
