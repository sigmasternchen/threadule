package models

import uuid "github.com/satori/go.uuid"

type Account struct {
	BaseModel
	UserID uuid.UUID `json:"-"`
	User   *User     `json:"-"`

	ScreenName    string `json:"screen_name"`
	TwitterHandle string `json:"twitter_handle"`
	TwitterID     *int64 `json:"twitter_id"`
	AvatarURL     string `json:"avatar_url"`

	RequestToken      *string `json:"-"`
	RequestSecret     *string `json:"-"`
	AccessToken       *string `json:"-"`
	AccessTokenSecret *string `json:"-"`
}
