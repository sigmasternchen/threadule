package models

type Group struct {
	BaseModel
	Users []*User `gorm:"many2many:user_groups;"`

	Name          string
	DisplayName   string
	AdminGroup    bool
	ManageUsers   bool
	ManageGroups  bool
	LimitAccounts uint
	LimitThreads  uint
	LimitTweets   uint
}

func GetDefaultAdminGroup() *Group {
	return &Group{
		Name:          "admin",
		DisplayName:   "Administrators",
		AdminGroup:    true,
		ManageUsers:   true,
		ManageGroups:  true,
		LimitAccounts: 0,
		LimitThreads:  0,
		LimitTweets:   0,
	}
}
