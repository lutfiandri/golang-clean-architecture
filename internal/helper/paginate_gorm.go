package helper

import (
	"math"

	"github.com/lutfiandri/golang-clean-architecture/internal/model"
	"gorm.io/gorm"
)

// how to use:
// db.Scopes(PaginateGorm(page, size)).Find()
func PaginateGorm(page, size *uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		p, s := getPageSize(page, size)
		offset := (p - 1) * s

		return db.Offset(offset).Limit(s)
	}
}

func GetPageMetadata(db *gorm.DB, entity any, page, size *uint) *model.PageMetadata {
	var itemCount int64
	db.Model(entity).Count(&itemCount)

	p, s := getPageSize(page, size)

	pageCount := int64(math.Ceil(float64(itemCount) / float64(s)))

	pageMetadata := model.PageMetadata{
		Page:      uint(p),
		Size:      uint(s),
		TotalItem: uint64(itemCount),
		TotalPage: uint64(pageCount),
	}

	return &pageMetadata
}

func getPageSize(page, size *uint) (int, int) {
	defaultPage := uint(1)
	defaultSize := uint(10)

	if page == nil {
		page = &defaultPage
	}

	if size == nil {
		size = &defaultSize
	}

	return int(*page), int(*size)
}
