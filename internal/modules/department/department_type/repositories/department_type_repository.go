package repositories

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
	"resedist/pkg/database"
	"resedist/pkg/pagination"

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
func (DepartmentTypeRepository *DepartmentTypeRepository) FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) []DepTypeModels.DepartmentType {
	var depTypes []DepTypeModels.DepartmentType
	var totalRows int64
	db := DepartmentTypeRepository.DB

	for _, scope := range scopes {
		db = db.Scopes(scope)
	}

	db.Model(&DepTypeModels.DepartmentType{}).Count(&totalRows)
	// fmt.Println(totalRows)
	pack.SetRows(totalRows)

	//fmt.Println(expand)
	//if expand {
	//	db = db.Preload("DepartmentType").Preload("Parent")
	//}

	result := db.Scopes(pack.ApplyToDB()).Find(&depTypes)
	if result.Error != nil {
		// Log the error or handle it as needed
		return nil
	}

	return depTypes
}
