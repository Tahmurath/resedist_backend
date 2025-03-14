package department_type

type ListDepartmentTypeRequest struct {
	Title          string `form:"title" json:"title"`
	DepartmentType string `form:"department_type" json:"department_type"`
	Sort           string `form:"sort" json:"sort"`
	Order          string `form:"order" json:"order"`
	Page           int    `form:"page" json:"page"`
	PageSize       int    `form:"page_size" json:"page_size"`
}
