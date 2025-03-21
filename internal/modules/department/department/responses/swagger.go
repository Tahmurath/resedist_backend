package responses

import "resedist/pkg/pagination"

type DepartmentResponse struct {
	ErrorCode string     `json:"_error_code" example:""`
	Status    string     `json:"_status" example:"success"`
	Data      Department `json:"data"`
	Message   string     `json:"message" example:"null"`
}
type DepartmentsResponse struct {
	ErrorCode  string `json:"_error_code" example:""`
	Status     string `json:"_status" example:"success"`
	Data       []Department
	Message    string              `json:"message" example:"null"`
	Pagination pagination.PagePack `json:"_pagination"`
}
