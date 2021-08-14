package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Session struct {
	BaseModel
	UserID uuid.UUID
	User   *User

	ValidUntil time.Time
}
