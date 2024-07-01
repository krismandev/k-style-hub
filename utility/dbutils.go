package utility

import (
	"k-style-test/model/request"

	"gorm.io/gorm"
)

func Paginate(dataParams map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// q := r.URL.Query()
		page := dataParams["page"].(int)
		if page <= 0 {
			page = 1
		}

		pageSize := dataParams["per_page"].(int)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Order(dataParams map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// q := r.URL.Query()

		if orderBy, ok := dataParams["order_by"]; ok && len(dataParams["order_by"].(string)) > 0 {
			if orderDir, dirOk := dataParams["order_dir"]; dirOk {
				dir := "ASC"
				if len(orderDir.(string)) > 0 {
					dir = orderDir.(string)
				}
				return db.Order(orderBy.(string) + " " + dir)
			}
			return db.Order(orderBy.(string))
		}

		return db
	}
}

func PreparePaginationAndOrderParam(structParam request.DataTableParam) map[string]interface{} {
	dataParams := make(map[string]interface{})

	dataParams["per_page"] = structParam.PerPage
	dataParams["page"] = structParam.Page
	dataParams["order_by"] = structParam.OrderBy
	dataParams["order_dir"] = structParam.OrderDir

	return dataParams

}
