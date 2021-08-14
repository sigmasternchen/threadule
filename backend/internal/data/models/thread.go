package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Thread struct {
	BaseModel
	AccountID uuid.UUID
	Account   *Account
	Tweets    []Tweet

	Sent         bool
	ScheduledFor time.Time
}
