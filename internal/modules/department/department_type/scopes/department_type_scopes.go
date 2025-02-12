package scopes

import (
	"gorm.io/gorm"
)

func TitleLike(title string) func(db *gorm.DB) *gorm.DB {
	if title != "" {
		return func(db *gorm.DB) *gorm.DB {

			query := "%" + title + "%"
			return db.Where("title LIKE ?", query)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func Sort(column string, order string) func(db *gorm.DB) *gorm.DB {

	var allowedSortColumns = map[string]bool{
		"id":         true,
		"title":      true,
		"created_at": true,
	}

	if order != "desc" {
		order = "asc"
	}

	if column != "" && allowedSortColumns[column] {
		return func(db *gorm.DB) *gorm.DB {

			return db.Order(column + " " + order)
		}
	}

	return func(db *gorm.DB) *gorm.DB {

		return db
	}
}
