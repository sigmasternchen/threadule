package data

import "threadule/backend/internal/data/models"

func (d *Data) UpdateTweet(tweet *models.Tweet) error {
	return d.db.
		Save(tweet).
		Error
}
