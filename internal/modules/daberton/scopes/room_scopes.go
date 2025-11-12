package scopes

import (
	"fmt"
	"strings"

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

func IdsOr(idParams string) func(db *gorm.DB) *gorm.DB {
	if idParams != "" {

		idStr := strings.Split(idParams, ",")
		var ids []uint

		for _, str := range idStr {
			var id uint
			fmt.Sscanf(str, "%d", &id)
			ids = append(ids, id)
		}

		return func(db *gorm.DB) *gorm.DB {

			return db.Or("id IN ?", ids)
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
