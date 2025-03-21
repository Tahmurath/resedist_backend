package department

type AddDepartmentRequest struct {
	Title            string `form:"title" json:"title" binding:"required,min=3,max=100"`
	DepartmentTypeId uint   `form:"departmenttypeid" json:"departmenttypeid" binding:"required"`
	ParentID         uint   `form:"parentid" json:"parentid" binding:"required"`
}
