package repositories

import (
	DepartmentModels "resedist/internal/modules/department/models"
	"resedist/pkg/database"

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

func (DepartmentRepository *DepartmentRepository) List(limit int) []DepartmentModels.Department {
	var department []DepartmentModels.Department

	DepartmentRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&department)

	return department
}

func (DepartmentRepository *DepartmentRepository) Find(id int) DepartmentModels.Department {
	var department DepartmentModels.Department

	DepartmentRepository.DB.Joins("User").First(&department, id)

	return department
}

func (DepartmentRepository *DepartmentRepository) Create(department DepartmentModels.Department) DepartmentModels.Department {
	var newDepartment DepartmentModels.Department

	DepartmentRepository.DB.Create(&department).Scan(&newDepartment)

	return newDepartment
}
