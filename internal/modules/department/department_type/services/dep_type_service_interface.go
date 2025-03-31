package services

import (
	DepTypeRequest "resedist/internal/modules/department/department_type/requests/deptype"
	DepTypeResponse "resedist/internal/modules/department/department_type/responses"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepTypeServiceInterface interface {
	Find(id uint) (DepTypeResponse.DepType, error)
	StoreAsUser(request DepTypeRequest.AddDepTypeRequest, user UserResponse.User) (DepTypeResponse.DepType, error)
	UpdateDepartment(id uint, request DepTypeRequest.EditDepTypeRequest, user UserResponse.User) (DepTypeResponse.DepType, error)
	Delete(id uint) error
	//UpdateAsUser(id int, request DepTypeRequest.EditDepartmentRequest, user UserResponse.User) (DepTypeResponse.Department, error)

	SearchDepTypesWithScopes(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) (DepTypeResponse.DepTypes, error)
	SearchDepTypesPaginated(request DepTypeRequest.ListDepTypeRequest) (DepTypeResponse.DepTypes, pagination.PagePack, error)
}
