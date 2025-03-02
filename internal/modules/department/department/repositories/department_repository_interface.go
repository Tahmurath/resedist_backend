package repositories

import (
	"gorm.io/gorm"
	DepartmentModels "resedist/internal/modules/department/department/models"
	"resedist/pkg/pagination"
)

type DepartmentRepositoryInterface interface {
	Find(id int, scopes ...func(*gorm.DB) *gorm.DB) DepartmentModels.Department
	Create(department DepartmentModels.Department) DepartmentModels.Department
	FindAllScope(pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) []DepartmentModels.Department
}
