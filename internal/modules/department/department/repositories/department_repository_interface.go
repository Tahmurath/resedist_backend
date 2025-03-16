package repositories

import (
	DepartmentModels "resedist/internal/modules/department/department/models"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentRepositoryInterface interface {
	Find(id int, scopes ...func(*gorm.DB) *gorm.DB) DepartmentModels.Department
	Create(department DepartmentModels.Department) DepartmentModels.Department
	FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) ([]DepartmentModels.Department, error)
}
