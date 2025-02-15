package responses

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
)

type DepartmentType struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type DepartmentTypes struct {
	Data []DepartmentType `json:"data"`
}

func ToDepartmentType(depType DepTypeModels.DepartmentType) DepartmentType {
	return DepartmentType{
		ID:    depType.ID,
		Title: depType.Title,
	}
}

func ToDepartmentTypes(depTypes []DepTypeModels.DepartmentType) DepartmentTypes {
	response := make([]DepartmentType, len(depTypes))

	for i, depType := range depTypes {
		response[i] = ToDepartmentType(depType)
	}

	return DepartmentTypes{Data: response}
}
