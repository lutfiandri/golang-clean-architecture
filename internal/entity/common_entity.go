package entity

import "time"

type BaseEntity struct {
	ID        uint      `gorm:"id;primaryKey"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
