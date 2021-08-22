package models

type User struct {
	BaseModel
	Groups []*Group `gorm:"many2many:user_groups;"`

	Username string `json:"username"`
	Password string `json:"password"`
}

func GetDefaultAdminUser() *User {
	return &User{
		Username: "admin",
	}
}
