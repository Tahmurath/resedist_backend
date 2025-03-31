package repositories

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepTypeRepositoryInterface interface {
	Find(id uint, scopes ...func(*gorm.DB) *gorm.DB) DepTypeModels.DepartmentType
	Create(depType DepTypeModels.DepartmentType) DepTypeModels.DepartmentType
	Delete(id uint) error
	Update(id uint, updates map[string]interface{}) (DepTypeModels.DepartmentType, error)
	FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) ([]DepTypeModels.DepartmentType, error)
}
