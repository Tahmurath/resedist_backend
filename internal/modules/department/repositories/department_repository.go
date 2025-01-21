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

func (DepartmentRepository *DepartmentRepository) FindAllByTitle(title string, limit int, expand bool) []DepartmentModels.Department {
	var departments []DepartmentModels.Department
	db := DepartmentRepository.DB
	if limit > 0 {
		db = db.Limit(limit)
	}
	if title != "" {
		query := "%" + title + "%"
		db = db.Where("title LIKE ?", query)
	}
	if expand {
		db = db.Preload("DepartmentType")
	}
	result := db.Find(&departments)
	if result.Error != nil {
		// Log the error or handle it as needed
		return nil
	}

	return departments
}

func (DepartmentRepository *DepartmentRepository) FindAll2(title string, limit int, expand bool) []DepartmentModels.Department {
	var departments []DepartmentModels.Department
	query := "%" + title + "%"

	db := DepartmentRepository.DB
	if limit > 0 {
		db = db.Limit(limit)
	}

	// Check if expand is true, and preload related data
	if expand {
		db = db.Preload("DepartmentType") // Assuming there's a DepartmentType relation
	}

	result := db.Where("title LIKE ?", query).Find(&departments)
	if result.Error != nil {
		// Log the error or handle it as needed
		return nil
	}

	return departments
}

// func (DepartmentRepository *DepartmentRepository) List(limit int) []DepartmentModels.Department {
// 	var department []DepartmentModels.Department

// 	DepartmentRepository.DB.Limit(limit).Joins("DepartmentType").Find(&department)

// 	return department
// }

// func (DepartmentRepository *DepartmentRepository) Find(id int) DepartmentModels.Department {
// 	var department DepartmentModels.Department

// 	DepartmentRepository.DB.Joins("DepartmentType").First(&department, id)

// 	return department
// }

func (DepartmentRepository *DepartmentRepository) Create(department DepartmentModels.Department) DepartmentModels.Department {
	var newDepartment DepartmentModels.Department

	DepartmentRepository.DB.Create(&department).Scan(&newDepartment)

	return newDepartment
}
