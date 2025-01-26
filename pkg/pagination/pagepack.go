package pagination

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"resedist/pkg/config"
	"strconv"
)

type PagePack struct {
	Page       int   `json:"page,omitempty;query:page"`
	PageSize   int   `json:"page_size,omitempty;query:page_size"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
	//Rows       interface{} `json:"rows"`
}

func New(c *gin.Context) *PagePack {
	cfg := config.Get()

	pp := NewConfig(c, cfg.URLKeys.Page, cfg.URLKeys.Pagesize)
	return pp
}
func NewConfig(c *gin.Context, pageKey, pageSizeKey string) *PagePack {
	page, _ := strconv.Atoi(c.DefaultQuery(pageKey, "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery(pageSizeKey, "10"))
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

		if p.Page <= 0 {
			p.Page = 1
		}

		switch {
		case p.PageSize > 30:
			p.PageSize = 30
		case p.PageSize <= 0:
			p.PageSize = 10
		}

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}
