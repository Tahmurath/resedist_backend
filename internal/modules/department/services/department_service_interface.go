package services

import (
	DepRequest "resedist/internal/modules/department/requests/department"
	DepResponse "resedist/internal/modules/department/responses"
	UserResponse "resedist/internal/modules/user/responses"
)

type DepartmentServiceInterface interface {
	Find(id int) (DepResponse.Department, error)
	List() DepResponse.Departments
	StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error)
}
