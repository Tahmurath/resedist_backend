package repositories

import (
	DepartmentModels "resedist/internal/modules/department/department/models"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentRepositoryInterface interface {
	Find(id uint, scopes ...func(*gorm.DB) *gorm.DB) DepartmentModels.Department
	Create(department DepartmentModels.Department) DepartmentModels.Department
	Delete(id uint) error
	Update(id uint, updates map[string]interface{}) (DepartmentModels.Department, error)
	FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) ([]DepartmentModels.Department, error)
}
