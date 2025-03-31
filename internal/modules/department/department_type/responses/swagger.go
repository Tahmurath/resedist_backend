package responses

import "resedist/pkg/pagination"

type NoContentResponse struct {
	ErrorCode string `json:"_error_code" example:""`
	Status    string `json:"_status" example:"success"`
	Message   string `json:"_message" example:"null"`
}
type DepTypeResponse struct {
	ErrorCode string  `json:"_error_code" example:""`
	Status    string  `json:"_status" example:"success"`
	Data      DepType `json:"data"`
	Message   string  `json:"_message" example:"null"`
}
type DepTypesResponse struct {
	ErrorCode  string              `json:"_error_code" example:""`
	Status     string              `json:"_status" example:"success"`
	Data       []DepType           `json:"data"`
	Message    string              `json:"_message" example:"null"`
	Pagination pagination.PagePack `json:"pagination"`
}
