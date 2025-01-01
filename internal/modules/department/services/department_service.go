package services

import (
	"errors"
	DepartmentModel "resedist/internal/modules/department/models"
	DepRepository "resedist/internal/modules/department/repositories"
	DepRequest "resedist/internal/modules/department/requests/department"
	DepResponse "resedist/internal/modules/department/responses"
	UserResponse "resedist/internal/modules/user/responses"
)

type DepartmentService struct {
	depRepository DepRepository.DepartmentRepositoryInterface
}

func New() *DepartmentService {
	return &DepartmentService{
		depRepository: DepRepository.New(),
	}
}

func (DepartmentService *DepartmentService) List() DepResponse.Departments {

	departments := DepartmentService.depRepository.List(4)
	return DepResponse.ToDepartments(departments)
}

func (DepartmentService *DepartmentService) Find(id int) (DepResponse.Department, error) {

	var response DepResponse.Department
	department := DepartmentService.depRepository.Find(id)

	if department.ID == 0 {
		return response, errors.New("department not found")
	}

	return DepResponse.ToDepartment(department), nil
}

func (DepartmentService *DepartmentService) StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error) {
	var department DepartmentModel.Department
	var response DepResponse.Department

	department.Title = request.Title
	department.DepartmentTypeId = request.DepartmentTypeId
	department.ParentID = request.ParentID
	department.AddedByUserID = user.ID

	newDepartment := DepartmentService.depRepository.Create(department)

	if newDepartment.ID == 0 {
		return response, errors.New("error in creating newDepartment")
	}

	return DepResponse.ToDepartment(newDepartment), nil
}
