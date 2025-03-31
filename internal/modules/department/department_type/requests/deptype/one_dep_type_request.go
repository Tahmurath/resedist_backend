package deptype

type OneDepTypeRequest struct {
	DepTypeId uint `uri:"id" binding:"required"`
}
