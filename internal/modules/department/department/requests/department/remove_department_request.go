package department

type RemoveDepartmentRequest struct {
	DepartmentId uint `uri:"id" binding:"required"`
}
