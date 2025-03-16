package repositories

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentTypeRepositoryInterface interface {
	// List(limit int) []DepartmentModels.Department
	// Find(id int) DepartmentModels.Department
	FindAll(title string, limit int) []DepTypeModels.DepartmentType
	FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) []DepTypeModels.DepartmentType
	//Create(depType DepartmentModels.DepartmentType) DepartmentModels.DepartmentType
}
