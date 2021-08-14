package models

import uuid "github.com/satori/go.uuid"

type Account struct {
	BaseModel
	UserID uuid.UUID
	User   *User

	ScreenName    string
	TwitterHandle string
	TwitterID     *int64
	AvatarURL     string

	RequestToken      *string
	RequestSecret     *string
	AccessToken       *string
	AccessTokenSecret *string
}
