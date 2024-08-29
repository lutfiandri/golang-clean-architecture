package entity

type User struct {
	BaseEntity
	RoleID        uint   `gorm:"role_id"`
	Name          string `gorm:"name"`
	Email         string `gorm:"email"`
	Password      string `gorm:"password"`
	Role          Role
	Organizations []Organization
}

func (e *User) TableName() string {
	return "users"
}
