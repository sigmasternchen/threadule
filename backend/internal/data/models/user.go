package models

type User struct {
	BaseModel
	Groups []*Group `gorm:"many2many:user_groups;"`

	Username string
	Password string
}
