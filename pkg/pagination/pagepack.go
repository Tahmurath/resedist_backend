package pagination

import (
	"math"

	"gorm.io/gorm"
)

type PagePack struct {
	Page       int   `json:"page,omitempty;query:page"`
	PageSize   int   `json:"page_size,omitempty;query:page_size"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}

func New(page, pageSize int) *PagePack {

	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize > 50:
		pageSize = 50
	case pageSize <= 0:
		pageSize = 10
	}

	pp := PagePack{
		Page:     page,
		PageSize: pageSize,
	}
	return &pp
}

func (p *PagePack) SetRows(rows int64) {
	p.TotalRows = rows

	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.PageSize)))
}

func (p *PagePack) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

// type Paginator interface {
// 	GetPage() int
// 	GetPageSize() int
// 	GetOffset() int
// 	ApplyToDB(db *gorm.DB) *gorm.DB
// 	//ApplyToQuery(query url.Values)
// }

// type PagePack struct {
// 	Page       int
// 	PageSize   int
// 	TotalRows  int64
// 	TotalPages int
// }

// func NewPagePack(page, pageSize int) *PagePack {
// 	if page <= 0 {
// 		page = 1
// 	}
// 	if pageSize > 50 {
// 		pageSize = 50
// 	} else if pageSize <= 0 {
// 		pageSize = 10
// 	}
// 	return &PagePack{Page: page, PageSize: pageSize}
// }

// func (p *PagePack) GetPage() int {
// 	return p.Page
// }

// func (p *PagePack) GetPageSize() int {
// 	return p.PageSize
// }

// func (p *PagePack) GetOffset() int {
// 	return (p.Page - 1) * p.PageSize
// }

// func (p *PagePack) ApplyToDB(db *gorm.DB) *gorm.DB {
// 	return db.Offset(p.GetOffset()).Limit(p.PageSize)
// }

// // func (p *PagePack) ApplyToQuery(query url.Values) {
// //     query.Set("page", strconv.Itoa(p.Page))
// //     query.Set("per_page", strconv.Itoa(p.PageSize))
// // }

// func (p *PagePack) SetRows(rows int64) {
// 	p.TotalRows = rows
// 	p.TotalPages = int(math.Ceil(float64(p.TotalRows) / float64(p.PageSize)))
// }
