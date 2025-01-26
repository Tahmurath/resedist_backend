package scopes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TitleLike(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	if title := c.Query("title"); title != "" {
		return func(db *gorm.DB) *gorm.DB {

			query := "%" + title + "%"
			return db.Where("title LIKE ?", query)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func ParentID(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	if parentid := c.Query("parentid"); parentid != "" {
		return func(db *gorm.DB) *gorm.DB {

			return db.Where("parent_id = ?", parentid)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
