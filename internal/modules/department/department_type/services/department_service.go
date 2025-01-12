package services

import (
	DepTypeRepository "resedist/internal/modules/department/department_type/repositories"
	DepTypeResponse "resedist/internal/modules/department/department_type/responses"
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
