package services

import (
	"errors"
	"gorm.io/gorm"
	DepartmentModel "resedist/internal/modules/department/models"
	DepRepository "resedist/internal/modules/department/repositories"
	DepRequest "resedist/internal/modules/department/requests/department"
	DepResponse "resedist/internal/modules/department/responses"
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

func (DepartmentService *DepartmentService) SearchP(pack *pagination.PagePack) DepResponse.Departments {

	departments := DepartmentService.depRepository.FindAllByTitleP(pack)

	return DepResponse.ToDepartments(departments, false)
}
func (DepartmentService *DepartmentService) SearchScope(expand bool, pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) DepResponse.Departments {

	departments := DepartmentService.depRepository.FindAllScope(expand, pack, scopes...)

	return DepResponse.ToDepartments(departments, expand)
}
func (DepartmentService *DepartmentService) Search(title string, page int, pageSize int, expand bool) DepResponse.Departments {

	departments := DepartmentService.depRepository.FindAllByTitle(title, page, pageSize, expand)

	return DepResponse.ToDepartments(departments, expand)
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

	return DepResponse.ToDepartment(newDepartment, true), nil
}

//func (DepartmentService *DepartmentService) SearchScope(expand bool, pack *pagination.PagePack, scopes ...func(*gorm.DB) *gorm.DB) DepResponse.Departments {
//
//	var updatedValues []func(*gorm.DB) *gorm.DB
//
//	if expand {
//		depModel := DepartmentModel.Department{}
//		updatedValues = append(scopes, DepScopes.Preload(expand, refl.GetFieldName(depModel, "DepartmentType"), refl.GetFieldName(depModel, "Parent")))
//	} else {
//		updatedValues = scopes
//	}
//
//	departments := DepartmentService.depRepository.FindAllScope(expand, pack, updatedValues...)
//
//	return DepResponse.ToDepartments(departments, expand)
//}
