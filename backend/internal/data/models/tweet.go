package models

import uuid "github.com/satori/go.uuid"

type Tweet struct {
	BaseModel
	ThreadID uuid.UUID

	TweetID *string
	Text    string
	Sent    bool
}
