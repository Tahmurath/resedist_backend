package repositories

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
	"resedist/pkg/database"

	"gorm.io/gorm"
)

type DepartmentTypeRepository struct {
	DB *gorm.DB
}

func New() *DepartmentTypeRepository {
	return &DepartmentTypeRepository{
		DB: database.Connection(),
	}
}

func (DepartmentTypeRepository *DepartmentTypeRepository) FindAll(title string, limit int) []DepTypeModels.DepartmentType {
	var depType []DepTypeModels.DepartmentType

	DepartmentTypeRepository.DB.Limit(limit).Where("title like ?", "%"+title+"%").Find(&depType)

	return depType
}
