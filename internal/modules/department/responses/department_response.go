package responses

import (
	departmentModels "resedist/internal/modules/department/models"
)

type Department struct {
	ID             uint
	Title          string
	DepartmentType uint
	Parent         uint
}

type Departments struct {
	Data []Department
}

func ToDepartment(department departmentModels.Department) Department {
	return Department{
		ID:             department.ID,
		Title:          department.Title,
		DepartmentType: department.DepartmentTypeId,
		Parent:         department.ParentID,
	}
}

func ToDepartments(departments []departmentModels.Department) Departments {
	var response Departments

	for _, department := range departments {
		response.Data = append(response.Data, ToDepartment(department))
	}

	return response
}
