package pagination

import (
	"math"

	"gorm.io/gorm"
)

type Paginator interface {
	GetPage() int
	GetPageSize() int
	GetOffset() int
	ApplyToDB() func(db *gorm.DB) *gorm.DB
	SetRows(rows int64)
	//ApplyToQuery(query url.Values)
}

type PagePack struct {
	Page       int   `json:"page,omitempty" query:"page"`
	PageSize   int   `json:"page_size,omitempty" query:"page_size"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`

	// 	Page       int   `json:"page,omitempty;query:page"`
	// 	PageSize   int   `json:"page_size,omitempty;query:page_size"`
	// 	TotalRows  int64 `json:"total_rows"`
	// 	TotalPages int   `json:"total_pages"`
}

func NewPagePack(page, pageSize int) *PagePack {
	if page <= 0 {
		page = 1
	}
	if pageSize > 50 {
		pageSize = 50
	} else if pageSize <= 0 {
		pageSize = 10
	}
	return &PagePack{Page: page, PageSize: pageSize}
}

func (p *PagePack) GetPage() int {
	return p.Page
}

func (p *PagePack) GetPageSize() int {
	return p.PageSize
}

func (p *PagePack) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *PagePack) ApplyToDB() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.PageSize)
	}
}

// func (p *PagePack) ApplyToQuery(query url.Values) {
//     query.Set("page", strconv.Itoa(p.Page))
//     query.Set("per_page", strconv.Itoa(p.PageSize))
// }

func (p *PagePack) SetRows(rows int64) {
	p.TotalRows = rows
	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.PageSize)))
}
