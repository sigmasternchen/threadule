package app

import "threadule/backend/internal/data/models"

type Data interface {
	CleanupSessions() error

	UpdateTweet(tweet *models.Tweet) error

	GetScheduledThreads() ([]models.Thread, error)
	GetTweetsForThread(thread *models.Thread) ([]models.Tweet, error)
	UpdateThread(thread *models.Thread) error
}
