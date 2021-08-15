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
	AccountID uuid.UUID `json:"-"`
	Account   *Account  `json:"account"`
	Tweets    []Tweet   `json:"tweets"`

	ScheduledFor time.Time    `json:"scheduled_for"`
	Status       ThreadStatus `json:"status"`
	Error        *string      `json:"error"`
}
