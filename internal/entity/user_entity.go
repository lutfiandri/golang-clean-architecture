package entity

type User struct {
	BaseEntity
	RoleID        uint `gorm:"role_id"`
	Role          Role
	Name          string          `gorm:"name"`
	Email         string          `gorm:"email"`
	Password      string          `gorm:"password"`
	Organizations []*Organization `gorm:"many2many:user_organizations"`
}

func (e *User) TableName() string {
	return "users"
}
