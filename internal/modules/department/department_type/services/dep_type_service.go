package services

import (
	"errors"
	"fmt"
	DepTypeModel "resedist/internal/modules/department/department_type/models"
	DepTypeRepository "resedist/internal/modules/department/department_type/repositories"
	DepTypeRequest "resedist/internal/modules/department/department_type/requests/deptype"
	DepTypeResponse "resedist/internal/modules/department/department_type/responses"
	DepTypeScopes "resedist/internal/modules/department/department_type/scopes"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepTypeService struct {
	depTypeRepo DepTypeRepository.DepTypeRepositoryInterface
}

func New() *DepTypeService {
	return &DepTypeService{
		depTypeRepo: DepTypeRepository.New(),
	}
}

func (s *DepTypeService) SearchDepTypesWithScopes(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) (DepTypeResponse.DepTypes, error) {
	depTypes, err := s.depTypeRepo.FindAllScope(pack, scopes...)
	if err != nil {
		return DepTypeResponse.DepTypes{}, err
	}
	return DepTypeResponse.ToDepTypes(depTypes), nil
}

func (s *DepTypeService) SearchDepTypesPaginated(request DepTypeRequest.ListDepTypeRequest) (DepTypeResponse.DepTypes, pagination.PagePack, error) {

	paginate := pagination.NewPagePack(request.Page, request.PageSize)

	scopes := []func(*gorm.DB) *gorm.DB{
		DepTypeScopes.TitleLike(request.Title),
		DepTypeScopes.IdsOr(request.DepType),
		DepTypeScopes.Activated(request.IsActive),
		DepTypeScopes.Sort(request.Sort, request.Order),
	}

	// departments, err := DepTypeService.SearchDepartmentsWithScopes(request.Expand, paginate, scopes...)
	// return departments, *paginate, err

	depTypes, err := s.depTypeRepo.FindAllScope(paginate, scopes...)
	if err != nil {
		return DepTypeResponse.DepTypes{}, *paginate, err
	}
	return DepTypeResponse.ToDepTypes(depTypes), *paginate, nil
}

func (s *DepTypeService) Find(id uint) (DepTypeResponse.DepType, error) {
	var response DepTypeResponse.DepType

	depType := s.depTypeRepo.Find(
		id,
	)

	if depType.ID == 0 {
		return response, errors.New("error in find Department")
	}

	return DepTypeResponse.ToDepType(depType), nil
}

func (s *DepTypeService) UpdateDepartment(id uint, request DepTypeRequest.EditDepTypeRequest, user UserResponse.User) (DepTypeResponse.DepType, error) {

	depType := s.depTypeRepo.Find(id)
	if depType.ID == 0 {
		return DepTypeResponse.DepType{}, errors.New("department not found")
	}

	updates := map[string]interface{}{
		"title":     request.Title,
		"is_active": request.IsActive,
	}

	updatedDepType, err := s.depTypeRepo.Update(id, updates)
	if err != nil {
		return DepTypeResponse.DepType{}, fmt.Errorf("failed to update department: %v", err)
	}

	return DepTypeResponse.ToDepType(updatedDepType), nil
}

func (s *DepTypeService) Delete(id uint) error {
	depType := s.depTypeRepo.Find(id)
	if depType.ID == 0 {
		return errors.New("depType not found")
	}

	return s.depTypeRepo.Delete(depType.ID)
}
func (s *DepTypeService) StoreAsUser(request DepTypeRequest.AddDepTypeRequest, user UserResponse.User) (DepTypeResponse.DepType, error) {
	var depType DepTypeModel.DepartmentType
	var response DepTypeResponse.DepType

	depType.Title = request.Title
	depType.AddedByUserID = user.ID
	depType.IsActive = request.IsActive

	newDepType := s.depTypeRepo.Create(depType)

	if newDepType.ID == 0 {
		return response, errors.New("error in creating newDepartment")
	}

	return DepTypeResponse.ToDepType(newDepType), nil
}
