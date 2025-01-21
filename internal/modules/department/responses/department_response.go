package responses

import (
	"resedist/internal/modules/department/department_type/responses"
	departmentModels "resedist/internal/modules/department/models"
)

type Department struct {
	ID             uint        `json:"id"`
	Title          string      `json:"title"`
	DepartmentType interface{} `json:"departmentType"`
	Parent         interface{} `json:"parent"`
}

type Departments struct {
	Data []Department `json:"data"`
}

type Parent struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	DepartmentType uint   `json:"departmentType"`
	Parent         uint   `json:"parent"`
}

func ToDepartment(department departmentModels.Department, expand bool) Department {
	response := Department{
		ID:    department.ID,
		Title: department.Title,
		//Parent: department.ParentID,
	}
	if expand && department.DepartmentType != nil {
		response.DepartmentType = responses.DepartmentType{
			ID:    department.DepartmentType.ID,
			Title: department.DepartmentType.Title,
		}
	} else {
		response.DepartmentType = department.DepartmentTypeId
	}

	if expand && department.Parent != nil {
		response.Parent = &Parent{
			ID:             department.Parent.ID,
			Title:          department.Parent.Title,
			DepartmentType: department.Parent.DepartmentTypeId,
			Parent:         department.Parent.ParentID,
		}
	} else if department.Parent != nil {
		response.Parent = &Department{ID: department.Parent.ID}
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
