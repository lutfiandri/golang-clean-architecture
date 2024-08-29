package helper

import (
	"math"

	"gorm.io/gorm"
)

// how to use:
// db.Scopes(PaginateGorm(page, size)).Find()
func PaginateGorm(page, size *uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		defaultPage := uint(1)
		defaultSize := uint(10)

		if page == nil {
			page = &defaultPage
		}

		if size == nil {
			size = &defaultSize
		}

		offset := (*page - 1) * *size

		return db.Offset(int(offset)).Limit(int(*size))
	}
}

func FindTotalItemAndPage(db *gorm.DB, entity any, page, size *uint) (uint64, uint64) {
	var itemCount int64
	db.Model(entity).Count(&itemCount)

	pageCount := int64(math.Ceil(float64(itemCount) / float64(*size)))

	return uint64(itemCount), uint64(pageCount)
}
