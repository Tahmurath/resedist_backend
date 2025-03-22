package department

type OneDepartmentRequest struct {
	DepartmentId uint `uri:"id" binding:"required"`
}
