package articles

type StoreRequest struct {
	Title   string `form:"title" json:"title"  binding:"required,min=3,max=100"`
	Content string `form:"content" json:"content"  binding:"required,min=3,max=60000"`
}
