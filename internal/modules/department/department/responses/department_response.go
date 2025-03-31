package responses

import (
	departmentModels "resedist/internal/modules/department/department/models"
	depaTypeModels "resedist/internal/modules/department/department_type/models"
	"resedist/internal/modules/department/department_type/responses"
)

type Department struct {
	ID             uint        `json:"id"`
	Title          string      `json:"title"`
	DepartmentType interface{} `json:"departmentType,omitempty"`
	Parent         interface{} `json:"parent,omitempty"`
	Label          string      `json:"label"`
}

type Departments struct {
	Data []Department `json:"data"`
}

// Helper function to convert department type
func mapDepartmentType(departmentType *depaTypeModels.DepartmentType) interface{} {
	if departmentType == nil {
		return nil
	}
	return responses.DepType{
		ID:    departmentType.ID,
		Title: departmentType.Title,
	}
}

// Helper function to convert parent
func mapParent(parent *departmentModels.Department) interface{} {
	if parent == nil {
		return nil
	}
	return Department{
		ID:             parent.ID,
		Title:          parent.Title,
		DepartmentType: parent.DepartmentTypeId,
		Parent:         parent.ParentID,
		Label:          "Bug",
	}
}

func ToDepartment(department departmentModels.Department, expand bool) Department {
	var departmentType interface{}
	var parent interface{}

	if expand {
		departmentType = mapDepartmentType(department.DepartmentType)
		parent = mapParent(department.Parent)
	} else {
		departmentType = department.DepartmentTypeId
		parent = department.ParentID
	}

	return Department{
		ID:             department.ID,
		Title:          department.Title,
		DepartmentType: departmentType,
		Parent:         parent,
		Label:          "Bug",
	}
}

func ToDepartments(departments []departmentModels.Department, expand bool) Departments {
	response := make([]Department, len(departments))

	for i, department := range departments {
		response[i] = ToDepartment(department, expand)
	}

	return Departments{Data: response}
}
