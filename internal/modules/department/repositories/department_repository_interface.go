package repositories

import (
	"gorm.io/gorm"
	DepartmentModels "resedist/internal/modules/department/models"
	"resedist/pkg/pagination"
)

type DepartmentRepositoryInterface interface {
	// List(limit int) []DepartmentModels.Department
	// Find(id int) DepartmentModels.Department
	Create(department DepartmentModels.Department) DepartmentModels.Department
	FindAllByTitle(title string, page int, pageSize int, expand bool) []DepartmentModels.Department
	FindAllByTitleP(pack *pagination.PagePack) []DepartmentModels.Department
	FindAllScope(expand bool, pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) []DepartmentModels.Department
}
