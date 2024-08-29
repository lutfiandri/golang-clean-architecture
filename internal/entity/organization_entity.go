package entity

type Organization struct {
	BaseEntity
	Name        string  `gorm:"name"`
	Description *string `gorm:"description"`
}

func (e *Organization) TableName() string {
	return "organizations"
}
