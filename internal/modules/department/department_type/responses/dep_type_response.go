package responses

import (
	DepTypeModels "resedist/internal/modules/department/department_type/models"
)

type DepType struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	IsActive bool   `json:"is_active"`
}

type DepTypes struct {
	Data []DepType `json:"data"`
}

func ToDepType(depType DepTypeModels.DepartmentType) DepType {
	return DepType{
		ID:       depType.ID,
		Title:    depType.Title,
		IsActive: depType.IsActive,
	}
}

func ToDepTypes(depTypes []DepTypeModels.DepartmentType) DepTypes {
	response := make([]DepType, len(depTypes))

	for i, depType := range depTypes {
		response[i] = ToDepType(depType)
	}

	return DepTypes{Data: response}
}
