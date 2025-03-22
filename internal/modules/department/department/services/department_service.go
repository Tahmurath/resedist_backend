package services

import (
	"errors"
	"fmt"
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
	depRepo DepRepository.DepartmentRepositoryInterface
}

func New() *DepartmentService {
	return &DepartmentService{
		depRepo: DepRepository.New(),
	}
}

func (s *DepartmentService) SearchDepartmentsWithScopes(expand bool, pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) (DepResponse.Departments, error) {
	departments, err := s.depRepo.FindAllScope(pack, scopes...)
	if err != nil {
		return DepResponse.Departments{}, err
	}
	return DepResponse.ToDepartments(departments, expand), nil
}

func (s *DepartmentService) SearchDepartmentsPaginated(request DepRequest.ListDepartmentRequest) (DepResponse.Departments, pagination.PagePack, error) {

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

	departments, err := s.depRepo.FindAllScope(paginate, scopes...)
	if err != nil {
		return DepResponse.Departments{}, *paginate, err
	}
	return DepResponse.ToDepartments(departments, request.Expand), *paginate, nil
}

func (s *DepartmentService) Find(id uint, expand bool, scopes ...func(*gorm.DB) *gorm.DB) (DepResponse.Department, error) {
	var response DepResponse.Department

	department := s.depRepo.Find(id, scopes...)

	if department.ID == 0 {
		return response, errors.New("error in find Department")
	}

	return DepResponse.ToDepartment(department, expand), nil
}

func (s *DepartmentService) UpdateDepartment(id uint, request DepRequest.EditDepartmentRequest, user UserResponse.User) (DepResponse.Department, error) {

	department := s.depRepo.Find(id)
	if department.ID == 0 {
		return DepResponse.Department{}, errors.New("department not found")
	}

	updates := map[string]interface{}{
		"title":              request.Title,
		"parent_id":          request.ParentID,
		"department_type_id": request.DepartmentTypeId,
	}

	updatedDepartment, err := s.depRepo.Update(id, updates)
	if err != nil {
		return DepResponse.Department{}, fmt.Errorf("failed to update department: %v", err)
	}

	return DepResponse.ToDepartment(updatedDepartment, false), nil
}

func (s *DepartmentService) Delete(id uint) error {
	department := s.depRepo.Find(id)
	if department.ID == 0 {
		return errors.New("department not found")
	}

	return s.depRepo.Delete(department.ID)
}
func (s *DepartmentService) StoreAsUser(request DepRequest.AddDepartmentRequest, user UserResponse.User) (DepResponse.Department, error) {
	var department DepartmentModel.Department
	var response DepResponse.Department

	department.Title = request.Title
	department.DepartmentTypeId = request.DepartmentTypeId
	department.ParentID = request.ParentID
	department.AddedByUserID = user.ID

	newDepartment := s.depRepo.Create(department)

	if newDepartment.ID == 0 {
		return response, errors.New("error in creating newDepartment")
	}

	return DepResponse.ToDepartment(newDepartment, true), nil
}
