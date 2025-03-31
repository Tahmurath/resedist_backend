package repositories

import (
	"errors"
	"fmt"
	DepTypeModels "resedist/internal/modules/department/department_type/models"
	"resedist/pkg/database"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepTypeRepository struct {
	DB *gorm.DB
}

func New() *DepTypeRepository {
	return &DepTypeRepository{
		DB: database.Connection(),
	}
}

func (DepTypeRepository *DepTypeRepository) Find(id uint, scopes ...func(*gorm.DB) *gorm.DB) DepTypeModels.DepartmentType {
	var depType DepTypeModels.DepartmentType
	db := DepTypeRepository.DB
	for _, scope := range scopes {
		db = db.Scopes(scope)
	}
	//db.Find(&department)
	db.First(&depType, id)

	return depType
}

func (r *DepTypeRepository) Delete(id uint) error {

	result := r.DB.Model(&DepTypeModels.DepartmentType{}).Delete("id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *DepTypeRepository) Update(id uint, updates map[string]interface{}) (DepTypeModels.DepartmentType, error) {
	result := r.DB.Model(&DepTypeModels.DepartmentType{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return DepTypeModels.DepartmentType{}, fmt.Errorf("failed to update department: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return DepTypeModels.DepartmentType{}, errors.New("department not found or no changes applied")
	}

	var updatedDepType DepTypeModels.DepartmentType
	if err := r.DB.First(&updatedDepType, id).Error; err != nil {
		return DepTypeModels.DepartmentType{}, fmt.Errorf("failed to fetch updated department: %v", err)
	}

	return updatedDepType, nil
}

func (DepTypeRepository *DepTypeRepository) Create(depType DepTypeModels.DepartmentType) DepTypeModels.DepartmentType {
	var newDepType DepTypeModels.DepartmentType

	DepTypeRepository.DB.Create(&depType).Scan(&newDepType)

	return newDepType
}

func (DepTypeRepository *DepTypeRepository) FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) ([]DepTypeModels.DepartmentType, error) {
	var depTypes []DepTypeModels.DepartmentType
	var totalRows int64
	db := DepTypeRepository.DB

	for _, scope := range scopes {
		db = db.Scopes(scope)
	}

	db.Model(&DepTypeModels.DepartmentType{}).Count(&totalRows)
	pack.SetRows(totalRows)

	result := db.Scopes(pack.ApplyToDB()).Find(&depTypes)
	if result.Error != nil {
		// Log the error or handle it as needed
		return nil, result.Error
	}

	return depTypes, nil
}
