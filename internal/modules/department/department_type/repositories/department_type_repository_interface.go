package repositories

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
)

type DepartmentTypeRepositoryInterface interface {
	// List(limit int) []DepartmentModels.Department
	// Find(id int) DepartmentModels.Department
	FindAll(title string, limit int) []DepTypeModels.DepartmentType
	//Create(depType DepartmentModels.DepartmentType) DepartmentModels.DepartmentType
}
