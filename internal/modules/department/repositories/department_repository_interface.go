package repositories

import (
	DepartmentModels "resedist/internal/modules/department/models"
)

type DepartmentRepositoryInterface interface {
	// List(limit int) []DepartmentModels.Department
	// Find(id int) DepartmentModels.Department
	Create(department DepartmentModels.Department) DepartmentModels.Department
	FindAllByTitle(title string, limit int, expand bool) []DepartmentModels.Department
}
