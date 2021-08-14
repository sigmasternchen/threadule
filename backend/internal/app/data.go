package app

import "threadule/backend/internal/data/models"

type Data interface {
	UpdateTweet(tweet *models.Tweet) error

	GetTweetsForThread(thread *models.Thread) ([]models.Tweet, error)
	UpdateThread(thread *models.Thread) error
}
