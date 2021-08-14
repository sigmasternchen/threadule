package models

type Group struct {
	BaseModel
	Users []*User `gorm:"many2many:user_groups;"`

	Name          string `json:"name"`
	DisplayName   string `json:"display_name"`
	AdminGroup    bool   `json:"admin_group"`
	ManageUsers   bool   `json:"manage_users"`
	ManageGroups  bool   `json:"manage_groups"`
	LimitAccounts uint   `json:"limit_accounts"`
	LimitThreads  uint   `json:"limit_threads"`
	LimitTweets   uint   `json:"limit_tweets"`
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
