package services

import (
	"errors"
	"gorm.io/gorm"
	DepartmentModel "resedist/internal/modules/department/department/models"
	DepRepository "resedist/internal/modules/department/department/repositories"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	DepResponse "resedist/internal/modules/department/department/responses"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"
)

type DepartmentService struct {
	depRepository DepRepository.DepartmentRepositoryInterface
}

func New() *DepartmentService {
	return &DepartmentService{
		depRepository: DepRepository.New(),
	}
}

func (DepartmentService *DepartmentService) SearchScope(expand bool, pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) DepResponse.Departments {

	departments := DepartmentService.depRepository.FindAllScope(pack, scopes...)

	return DepResponse.ToDepartments(departments, expand)
}

func (DepartmentService *DepartmentService) Find(id int, expand bool, scopes ...func(*gorm.DB) *gorm.DB) (DepResponse.Department, error) {
	var response DepResponse.Department

	department := DepartmentService.depRepository.Find(id, scopes...)

	if department.ID == 0 {
		return response, errors.New("error in find Department")
	}

	return DepResponse.ToDepartment(department, expand), nil
}

//func (DepartmentService *DepartmentService) Search(title string, page int, pageSize int, expand bool) DepResponse.Departments {
//
//	departments := DepartmentService.depRepository.FindAllByTitle(title, page, pageSize, expand)
//
//	return DepResponse.ToDepartments(departments, expand)
//}

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

	return DepResponse.ToDepartment(newDepartment, true), nil
}
