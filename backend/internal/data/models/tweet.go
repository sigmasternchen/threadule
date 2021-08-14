package models

import uuid "github.com/satori/go.uuid"

type TweetStatus string

const (
	TweetScheduled TweetStatus = "SCHEDULED"
	TweetFailed    TweetStatus = "FAILED"
	TweetDone      TweetStatus = "DONE"
)

type Tweet struct {
	BaseModel
	ThreadID uuid.UUID

	Text    string
	Ordinal int
	Status  TweetStatus
	TweetID *int64
	Error   *string
}
