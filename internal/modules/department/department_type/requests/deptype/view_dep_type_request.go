package deptype

type ListDepTypeRequest struct {
	Title    string `form:"title" json:"title"`
	DepType  string `form:"depType" json:"depType"`
	Expand   bool   `form:"expand" json:"expand"`
	IsActive bool   `form:"is_active" json:"is_active"`
	Sort     string `form:"sort" json:"sort"`
	Order    string `form:"order" json:"order"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
}

type ShowDepTypeRequest struct {
	Expand bool `form:"expand"`
}
