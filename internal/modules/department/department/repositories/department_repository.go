package repositories

import (
	"errors"
	"fmt"
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

func (DepartmentRepository *DepartmentRepository) Find(id uint, scopes ...func(*gorm.DB) *gorm.DB) DepartmentModels.Department {
	var department DepartmentModels.Department
	db := DepartmentRepository.DB
	for _, scope := range scopes {
		db = db.Scopes(scope)
	}
	//db.Find(&department)
	db.First(&department, id)

	return department
}

func (r *DepartmentRepository) Delete(id uint) error {

	result := r.DB.Model(&DepartmentModels.Department{}).Delete("id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *DepartmentRepository) Update(id uint, updates map[string]interface{}) (DepartmentModels.Department, error) {
	result := r.DB.Model(&DepartmentModels.Department{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return DepartmentModels.Department{}, fmt.Errorf("failed to update department: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return DepartmentModels.Department{}, errors.New("department not found or no changes applied")
	}

	var updatedDepartment DepartmentModels.Department
	if err := r.DB.First(&updatedDepartment, id).Error; err != nil {
		return DepartmentModels.Department{}, fmt.Errorf("failed to fetch updated department: %v", err)
	}

	return updatedDepartment, nil
}

func (DepartmentRepository *DepartmentRepository) Create(department DepartmentModels.Department) DepartmentModels.Department {
	var newDepartment DepartmentModels.Department

	DepartmentRepository.DB.Create(&department).Scan(&newDepartment)

	return newDepartment
}

func (DepartmentRepository *DepartmentRepository) FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) ([]DepartmentModels.Department, error) {
	var departments []DepartmentModels.Department
	var totalRows int64
	db := DepartmentRepository.DB

	for _, scope := range scopes {
		db = db.Scopes(scope)
	}

	db.Model(&DepartmentModels.Department{}).Count(&totalRows)
	pack.SetRows(totalRows)

	result := db.Scopes(pack.ApplyToDB()).Find(&departments)
	if result.Error != nil {
		// Log the error or handle it as needed
		return nil, result.Error
	}

	return departments, nil
}
