package responses

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
)

type DepartmentType struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type DepartmentTypes struct {
	Data []DepartmentType
}

func ToDepartmentType(depType DepTypeModels.DepartmentType) DepartmentType {
	return DepartmentType{
		ID:    depType.ID,
		Title: depType.Title,
	}
}

func ToDepartmentTypes(depTypes []DepTypeModels.DepartmentType) DepartmentTypes {
	var response DepartmentTypes

	for _, depType := range depTypes {
		response.Data = append(response.Data, ToDepartmentType(depType))
	}

	return response
}
