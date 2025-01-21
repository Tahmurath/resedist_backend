package responses

import (
	"resedist/internal/modules/department/department_type/responses"
	departmentModels "resedist/internal/modules/department/models"
)

type Department struct {
	ID             uint        `json:"id"`
	Title          string      `json:"title"`
	DepartmentType interface{} `json:"departmentType"`
	Parent         uint        `json:"parent"`
}

type Departments struct {
	Data []Department `json:"data"`
}

func ToDepartment(department departmentModels.Department, expand bool) Department {
	response := Department{
		ID:     department.ID,
		Title:  department.Title,
		Parent: department.ParentID,
	}
	if expand && department.DepartmentType != nil {
		response.DepartmentType = responses.DepartmentType{
			ID:    department.DepartmentType.ID,
			Title: department.DepartmentType.Title,
		}
	} else {
		response.DepartmentType = department.DepartmentTypeId
	}
	return response
}

func ToDepartments(departments []departmentModels.Department, expand bool) Departments {
	var response Departments

	for _, department := range departments {
		response.Data = append(response.Data, ToDepartment(department, expand))
	}

	return response
}
