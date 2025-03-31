package deptype

type EditDepTypeRequest struct {
	Title    string `form:"title" json:"title" binding:"required,min=3,max=100"`
	IsActive bool   `form:"is_active" json:"is_active"`
}
