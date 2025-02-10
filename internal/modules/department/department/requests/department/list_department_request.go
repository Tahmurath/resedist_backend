package department

type ListDepartmentRequest struct {
	Title            string `form:"title" json:"title"`
	DepartmentTypeId int    `form:"departmenttypeid" json:"departmenttypeid"`
	ParentID         int    `form:"parentid" json:"parentid"`
	Expand           bool   `form:"expand" json:"expand"`
	Sort             string `form:"sort" json:"sort"`
	Order            string `form:"order" json:"order"`
	Page             int    `form:"page" json:"page"`
	PageSize         int    `form:"page_size" json:"page_size"`
}
