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

func Activated(isactive bool) func(db *gorm.DB) *gorm.DB {
	if isactive {
		return func(db *gorm.DB) *gorm.DB {

			return db.Where("is_active = ?", isactive)
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

func ParentIDS(parentIdParams string) func(db *gorm.DB) *gorm.DB {
	if parentIdParams != "" {

		parentIdStr := strings.Split(parentIdParams, ",")
		var parents []uint

		for _, str := range parentIdStr {
			var parent uint
			//fmt.Sscanf(str, "%d", &parent)
			if _, err := fmt.Sscanf(str, "%d", &parent); err != nil {
				continue
			}
			parents = append(parents, parent)
		}

		return func(db *gorm.DB) *gorm.DB {

			return db.Where("parent_id IN ?", parents)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func DepTypes(deptypeParams string) func(db *gorm.DB) *gorm.DB {
	if deptypeParams != "" {

		parentIdStr := strings.Split(deptypeParams, ",")
		var deptypes []int

		for _, str := range parentIdStr {
			var deptype int
			fmt.Sscanf(str, "%d", &deptype)
			deptypes = append(deptypes, deptype)
		}

		return func(db *gorm.DB) *gorm.DB {

			return db.Where("department_type_id IN ?", deptypes)
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

		// for _, str := range idStr {
		// 	id, err := strconv.Atoi(strings.TrimSpace(str))
		// 	if err != nil {
		// 		continue // یا خطا رو برگردون
		// 	}
		// 	ids = append(ids, id)
		// }

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
		"id":    true,
		"title": true,
		//"department_type_id": true,
		//"parent_id ":         true,
		"created_at": true,
		"is_active":  true,
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
