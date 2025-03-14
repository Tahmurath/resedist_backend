package department

type ListDepartmentRequest struct {
	Title            string `form:"title" json:"title"`
	Department       string `form:"department" json:"department"`
	DepartmentTypeId int    `form:"departmenttypeid" json:"departmenttypeid"`
	DepartmentType   string `form:"department_type" json:"department_type"`
	ParentID         int    `form:"parentid" json:"parentid"`
	Parent           string `form:"parent" json:"parent"`
	Expand           bool   `form:"expand" json:"expand"`
	Sort             string `form:"sort" json:"sort"`
	Order            string `form:"order" json:"order"`
	Page             int    `form:"page" json:"page"`
	PageSize         int    `form:"page_size" json:"page_size"`
}

type OneDepartmentRequest struct {
	DepartmentId int    `uri:"id"`
	Department   string `form:"department" json:"department"`
	Expand       bool   `form:"expand"`
}
