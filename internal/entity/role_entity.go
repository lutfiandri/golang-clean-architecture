package entity

type Role struct {
	BaseEntity
	Name string `gorm:"name"`
}

func (e *Role) TableName() string {
	return "roles"
}
