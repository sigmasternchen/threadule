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
	ThreadID uuid.UUID `json:"-"`

	Text    string      `json:"text"`
	Ordinal int         `json:"ordinal"`
	Status  TweetStatus `json:"status"`
	TweetID *int64      `json:"tweet_id"`
	Error   *string     `json:"error"`
}
