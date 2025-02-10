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

func Preload(Expand bool, preloades ...string) func(db *gorm.DB) *gorm.DB {
	if Expand && len(preloades) > 0 {
		return func(db *gorm.DB) *gorm.DB {

			for _, scope := range preloades {
				db = db.Preload(scope)
			}
			return db
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func ParentID(parentId int) func(db *gorm.DB) *gorm.DB {
	if parentId > 0 {
		return func(db *gorm.DB) *gorm.DB {

			return db.Where("parent_id = ?", parentId)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func Sort(column string, order string) func(db *gorm.DB) *gorm.DB {

	var allowedSortColumns = map[string]bool{
		"id":                 true,
		"title":              true,
		"department_type_id": true,
		"parent_id ":         true,
		"created_at":         true,
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
