package services

import (
	DepTypeRepository "resedist/internal/modules/department/department_type/repositories"
	DepTypeResponse "resedist/internal/modules/department/department_type/responses"
	"resedist/pkg/pagination"

	"gorm.io/gorm"
)

type DepartmentTypeService struct {
	depTypeRepository DepTypeRepository.DepartmentTypeRepositoryInterface
}

func New() *DepartmentTypeService {
	return &DepartmentTypeService{
		depTypeRepository: DepTypeRepository.New(),
	}
}

func (DepartmentTypeService *DepartmentTypeService) Search(title string) DepTypeResponse.DepartmentTypes {

	depTypes := DepartmentTypeService.depTypeRepository.FindAll(title, 10)

	return DepTypeResponse.ToDepartmentTypes(depTypes)
}

func (DepartmentTypeService *DepartmentTypeService) SearchScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) DepTypeResponse.DepartmentTypes {

	depTypes := DepartmentTypeService.depTypeRepository.FindAllScope(pack, scopes...)

	return DepTypeResponse.ToDepartmentTypes(depTypes)
}
