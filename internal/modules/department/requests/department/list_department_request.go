package department

type ListDepartmentRequest struct {
	Title            string `form:"title" json:"title" binding:"min=3,max=100"`
	DepartmentTypeId uint   `form:"departmenttypeid" json:"departmenttypeid"`
	ParentID         uint   `form:"parentid" json:"parentid"`
}
