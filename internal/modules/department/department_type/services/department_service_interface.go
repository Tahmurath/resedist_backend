package services

import (
	DepTypeResponse "resedist/internal/modules/department/department_type/responses"
)

type DepartmentTypeServiceInterface interface {
	//Find(id int) (DepResponse.Department, error)
	Search(title string) DepTypeResponse.DepartmentTypes
	// StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)
}
