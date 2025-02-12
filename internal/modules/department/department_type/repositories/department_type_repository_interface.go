package repositories

import (
	"gorm.io/gorm"
	DepTypeModels "resedist/internal/modules/department/department_type/models"
	"resedist/pkg/pagination"
)

type DepartmentTypeRepositoryInterface interface {
	// List(limit int) []DepartmentModels.Department
	// Find(id int) DepartmentModels.Department
	FindAll(title string, limit int) []DepTypeModels.DepartmentType
	FindAllScope(pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) []DepTypeModels.DepartmentType
	//Create(depType DepartmentModels.DepartmentType) DepartmentModels.DepartmentType
}
