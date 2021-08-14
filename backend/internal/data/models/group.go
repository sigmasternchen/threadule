package models

type Group struct {
	BaseModel
	Users []*User `gorm:"many2many:user_groups;"`

	Name        string
	DisplayName string

	LimitAccounts uint
	LimitThreads  uint
	LimitTweets   uint
}
