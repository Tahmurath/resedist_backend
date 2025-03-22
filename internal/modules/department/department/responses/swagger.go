package responses

import "resedist/pkg/pagination"

type NoContentResponse struct {
	ErrorCode string `json:"_error_code" example:""`
	Status    string `json:"_status" example:"success"`
	Message   string `json:"_message" example:"null"`
}
type DepartmentResponse struct {
	ErrorCode string     `json:"_error_code" example:""`
	Status    string     `json:"_status" example:"success"`
	Data      Department `json:"data"`
	Message   string     `json:"_message" example:"null"`
}
type DepartmentsResponse struct {
	ErrorCode  string              `json:"_error_code" example:""`
	Status     string              `json:"_status" example:"success"`
	Data       []Department        `json:"data"`
	Message    string              `json:"_message" example:"null"`
	Pagination pagination.PagePack `json:"pagination"`
}
