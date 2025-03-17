package services

import (
	"errors"
	DepartmentModel "resedist/internal/modules/department/department/models"
	DepRepository "resedist/internal/modules/department/department/repositories"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	DepResponse "resedist/internal/modules/department/department/responses"
	DepScopes "resedist/internal/modules/department/department/scopes"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentService struct {
	depRepository DepRepository.DepartmentRepositoryInterface
}

func New() *DepartmentService {
	return &DepartmentService{
		depRepository: DepRepository.New(),
	}
}

func (DepartmentService *DepartmentService) SearchDepartmentsWithScopes(expand bool, pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) (DepResponse.Departments, error) {
	departments, err := DepartmentService.depRepository.FindAllScope(pack, scopes...)
	if err != nil {
		return DepResponse.Departments{}, err
	}
	return DepResponse.ToDepartments(departments, expand), nil
}

func (DepartmentService *DepartmentService) SearchDepartmentsPaginated(request DepRequest.ListDepartmentRequest) (DepResponse.Departments, pagination.PagePack, error) {

	paginate := pagination.NewPagePack(request.Page, request.PageSize)

	scopes := []func(*gorm.DB) *gorm.DB{
		DepScopes.TitleLike(request.Title),
		DepScopes.IdsOr(request.Department),
		DepScopes.Preload(request.Expand, "DepartmentType", "Parent"),
		DepScopes.ParentIDS(request.Parent),
		DepScopes.DepTypes(request.DepartmentType),
		DepScopes.Sort(request.Sort, request.Order),
	}

	// departments, err := DepartmentService.SearchDepartmentsWithScopes(request.Expand, paginate, scopes...)
	// return departments, *paginate, err

	departments, err := DepartmentService.depRepository.FindAllScope(paginate, scopes...)
	if err != nil {
		return DepResponse.Departments{}, *paginate, err
	}
	return DepResponse.ToDepartments(departments, request.Expand), *paginate, nil
}

func (DepartmentService *DepartmentService) Find(id int, expand bool, scopes ...func(*gorm.DB) *gorm.DB) (DepResponse.Department, error) {
	var response DepResponse.Department

	department := DepartmentService.depRepository.Find(id, scopes...)

	if department.ID == 0 {
		return response, errors.New("error in find Department")
	}

	return DepResponse.ToDepartment(department, expand), nil
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
