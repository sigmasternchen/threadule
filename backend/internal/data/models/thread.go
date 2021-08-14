package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type ThreadStatus string

const (
	ThreadScheduled  ThreadStatus = "SCHEDULED"
	ThreadProcessing ThreadStatus = "PROCESSING"
	ThreadFailed     ThreadStatus = "FAILED"
	ThreadDone       ThreadStatus = "DONE"
)

type Thread struct {
	BaseModel
	AccountID uuid.UUID
	Account   *Account
	Tweets    []Tweet

	ScheduledFor time.Time
	Status       ThreadStatus
	Error        *string
}
