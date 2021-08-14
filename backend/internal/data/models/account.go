package models

import uuid "github.com/satori/go.uuid"

type Account struct {
	BaseModel
	UserID uuid.UUID
	User   *User

	Name              string
	AccessToken       *string
	AccessTokenSecret *string
}
