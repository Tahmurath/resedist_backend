package repositories

import (
	DepartmentModels "resedist/internal/modules/department/department/models"
	"resedist/pkg/database"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	DB *gorm.DB
}

func New() *DepartmentRepository {
	return &DepartmentRepository{
		DB: database.Connection(),
	}
}

func (DepartmentRepository *DepartmentRepository) Find(id int, scopes ...func(*gorm.DB) *gorm.DB) DepartmentModels.Department {
	var department DepartmentModels.Department
	db := DepartmentRepository.DB
	for _, scope := range scopes {
		db = db.Scopes(scope)
	}
	//db.Find(&department)
	db.First(&department, id)

	return department
}

func (DepartmentRepository *DepartmentRepository) Create(department DepartmentModels.Department) DepartmentModels.Department {
	var newDepartment DepartmentModels.Department

	DepartmentRepository.DB.Create(&department).Scan(&newDepartment)

	return newDepartment
}

func (DepartmentRepository *DepartmentRepository) FindAllScope(pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) []DepartmentModels.Department {
	var departments []DepartmentModels.Department
	var totalRows int64
	db := DepartmentRepository.DB

	for _, scope := range scopes {
		db = db.Scopes(scope)
	}

	db.Model(&DepartmentModels.Department{}).Count(&totalRows)
	pack.SetRows(totalRows)

	result := db.Scopes(pack.Paginate()).Find(&departments)
	if result.Error != nil {
		// Log the error or handle it as needed
		return nil
	}

	return departments
}
